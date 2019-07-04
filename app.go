package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-redis/redis"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/atpb"
	"github.com/imulab-z/access-token-service/pkg"
	"github.com/imulab-z/access-token-service/pkg/mw"
	"github.com/imulab-z/access-token-service/pkg/svc"
	grpctransport "github.com/imulab-z/access-token-service/pkg/transport/grpc"
	httptransport "github.com/imulab-z/access-token-service/pkg/transport/http"
	discoverysvc "github.com/imulab-z/discovery-service/exported"
	keysvc "github.com/imulab-z/key-service/exported"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/square/go-jose.v2"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// setup logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var (
		app                 *cli.App
		argHttpPort         int
		argGrpcPort         int
		argRedisHost        string
		argRedisPort        int
		argRedisPwd         string
		argRedisDb          int
		argTokenStrategy    string
		argSigningAlg       string
		argDefaultTtl       int64
		argDefaultLeeway    int64
		argKeySvcHost       string
		argKeySvcPort       int
		argDiscoverySvcHost string
		argDiscoverySvcPort int
	)
	{
		app = cli.NewApp()
		app.Name = "access-token-service"
		app.Flags = []cli.Flag{
			cli.IntFlag{
				Name:        "http-port, p",
				Value:       8080,
				EnvVar:      "HTTP_PORT",
				Destination: &argHttpPort,
				Usage:       "Port the http service listens on.",
			},
			cli.IntFlag{
				Name:        "grpc-port, q",
				Value:       8081,
				EnvVar:      "GRPC_PORT",
				Destination: &argGrpcPort,
				Usage:       "Port the grpc service listens on.",
			},
			cli.StringFlag{
				Name:        "redis-host",
				Value:       "localhost",
				EnvVar:      "REDIS_HOST",
				Destination: &argRedisHost,
				Usage:       "Host of the redis database",
			},
			cli.IntFlag{
				Name:        "redis-port",
				Value:       6379,
				EnvVar:      "REDIS_PORT",
				Destination: &argRedisPort,
				Usage:       "Port of the redis database",
			},
			cli.StringFlag{
				Name:        "redis-pwd",
				Value:       "",
				EnvVar:      "REDIS_PWD",
				Destination: &argRedisPwd,
				Usage:       "Password of the redis database",
			},
			cli.IntFlag{
				Name:        "redis-db",
				Value:       0,
				EnvVar:      "REDIS_DB",
				Destination: &argRedisDb,
				Usage:       "DB number of the redis database",
			},
			cli.StringFlag{
				Name:        "token-strategy",
				Value:       pkg.TokenStrategyJwt,
				EnvVar:      "TOKEN_STRATEGY",
				Destination: &argTokenStrategy,
				Usage:       "Algorithm to use for generating the refresh token",
			},
			cli.StringFlag{
				Name:        "signing-algorithm",
				Value:       "RS256",
				EnvVar:      "SIGNING_ALGORITHM",
				Destination: &argSigningAlg,
				Usage:       "Algorithm to use when signing the access token.",
			},
			cli.Int64Flag{
				Name:        "default-ttl",
				Value:       int64(time.Hour / time.Second),
				EnvVar:      "DEFAULT_TTL",
				Destination: &argDefaultTtl,
				Usage:       "Number of seconds until the token expire.",
			},
			cli.Int64Flag{
				Name:        "default-leeway",
				Value:       int64(2 * time.Minute / time.Second),
				EnvVar:      "DEFAULT_LEEWAY",
				Destination: &argDefaultLeeway,
				Usage:       "Number of seconds in clock skew to tolerate during token validation.",
			},
			cli.StringFlag{
				Name:        "key-service-host",
				Value:       "key-service",
				EnvVar:      "KEY_SERVICE_HOST",
				Destination: &argKeySvcHost,
				Usage:       "Hostname of the key service",
			},
			cli.IntFlag{
				Name:        "key-service-port",
				Value:       80,
				EnvVar:      "KEY_SERVICE_PORT",
				Destination: &argKeySvcPort,
				Usage:       "Port of the key service",
			},
			cli.StringFlag{
				Name:        "discovery-service-host",
				Value:       "discovery-service",
				EnvVar:      "DISCOVERY_SERVICE_HOST",
				Destination: &argDiscoverySvcHost,
				Usage:       "Hostname of the discovery service",
			},
			cli.IntFlag{
				Name:        "discovery-service-port",
				Value:       80,
				EnvVar:      "DISCOVERY_SERVICE_PORT",
				Destination: &argDiscoverySvcPort,
				Usage:       "Port of the discovery service",
			},
		}
		app.Action = func(c *cli.Context) error {
			var errors chan error
			{
				errors = make(chan error)
				defer close(errors)
			}

			var redisClient *redis.Client
			{
				redisClient = redis.NewClient(&redis.Options{
					Addr:     fmt.Sprintf("%s:%d", argRedisHost, argRedisPort),
					Password: argRedisPwd,
					DB:       argRedisDb,
				})
			}

			var discovery *discoverysvc.Discovery
			{
				discovery, _ = discoverysvc.NewDiscoveryClient(
					argDiscoverySvcHost,
					argDiscoverySvcPort,
					logger,
				).Get(context.Background())
				if discovery == nil {
					return pkg.ErrServer("failed to obtain discovery")
				}
				logger.Log("discovery", "resolved")
			}

			var keySet *jose.JSONWebKeySet
			{
				keySet = keysvc.NewKeyClient(
					argKeySvcHost,
					argKeySvcPort,
					logger,
				).GetJsonWebKeySet(keysvc.IncludePrivate)
				if keySet == nil {
					return pkg.ErrServer("failed to obtain key set")
				}
				logger.Log("key_set", "resolved")
			}

			var strategy pkg.TokenStrategy
			{
				switch argTokenStrategy {
				case pkg.TokenStrategyJwt:
					strategy = pkg.NewJwtTokenStrategy(
						keySet,
						jose.SignatureAlgorithm(argSigningAlg),
						discovery.Issuer,
						argDefaultTtl,
						argDefaultLeeway,
						logger,
					)
				default:
					return pkg.ErrServer("invalid token strategy")
				}
			}

			var service exported.Service
			{
				service = svc.NewService(redisClient, strategy, argDefaultTtl, argDefaultLeeway, logger)
				service = mw.NewLoggingMiddleware(service, logger)
			}

			go func() {
				c := make(chan os.Signal)
				signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
				errors <- fmt.Errorf("%s", <-c)
			}()

			go func() {
				server := httptransport.NewHTTPServer(redisClient, service, logger)
				addr := fmt.Sprintf(":%d", argHttpPort)
				logger.Log("transport", "HTTP", "addr", addr)
				errors <- http.ListenAndServe(addr, server)
			}()

			go func() {
				addr := fmt.Sprintf(":%d", argGrpcPort)
				grpcListener, err := net.Listen("tcp", addr)
				if err != nil {
					errors <- err
					os.Exit(1)
				}
				logger.Log("transport", "gRPC", "addr", addr)

				baseServer := grpc.NewServer(grpc.UnaryInterceptor(gt.Interceptor))
				reflection.Register(baseServer)

				server := grpctransport.NewGrpcServer(service, logger)

				atpb.RegisterAccessTokenServiceServer(baseServer, server)
				errors <- baseServer.Serve(grpcListener)
			}()

			return <-errors
		}
	}

	logger.Log("fatal", app.Run(os.Args))
}
