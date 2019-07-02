package main

import (
	"context"
	keysvc "github.com/imulab-z/key-service/exported"
)
import discoverysvc "github.com/imulab-z/discovery-service/exported"

func main() {
	keysvc.NewKeyClient("", 0, nil).GetJsonWebKeySet(keysvc.IncludePrivate)
	discoverysvc.NewDiscoveryClient("", 0, nil).Get(context.Background())
}
