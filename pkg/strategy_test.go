package pkg

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/square/go-jose.v2"
	"strings"
	"testing"
)

var (
	testJwks = `
{
   "keys":[
      {
         "kty":"RSA",
         "kid":"9e679316-8092-47b0-b857-c04e1a900a2c",
         "use":"sig",
         "alg":"RS256",
         "n":"zbnENxDy1h9MMLHedpjDGfriFsUxNjRir_HPhLlBG5lVoeVD36oa-UQzxeybn60rXhU25mmH8mype9IOug7iGmPttNk5N_6gpujqlXyI9D_9gZc_Q7q86g_2-bkmwYBBMvO44_xxUyxq2qtXnOjEcmQrGNa8Bma3JmLn8g7kAVUadIvLVd314VHqd5zlQWCm9FgJkHwLwM5wsoL82oPzvfO_Gp3jGA7PxZt-_dKxseSmzMt2cgAa_R1APvPTmvWqZBmL0ofIrjmGwCBnDbaC1cfVW1ctBL8p-BS9Zg8UP-TbI6YE-z55npTj625jQnY2z98uu03lFKxUPYGOU3t3Rw",
         "e":"AQAB",
         "d":"KWK1ll5Se8DQEvu1RSZ2GUPfN7VzsPPY5ow-svSRpHu0Pl1gyh3uTzbpM2fl5rlvf_3EjZWtQ5eYgNBBJizYT3bK2xbX6-vNZcZ1ap0i7__vf6Jtl7J5TpznklUs9sBWXWmqSDMAmTrWRCcz-NzaqCh-gwCv0pnXPmGnR4q3U4zugOdDLJEB6B60NgEaAmWHcqLnIPicbPZlXZ6_IuzgR4JCSLv8rJ4UEtDeU_6BqbFt5mw3kbFAE2sEbEKF4bdTtcwSpIKfsoKVIoI3fWR77vl3xIYCd4RD8wWPvHZblpA0wc3sezZNdBLP749Ev_emR3fD83r1LYp8QxtusCBgAQ",
         "p":"_NjeksH49SHkmlhlumtzFjP5my_s0oRWTRnb5lEG0iulmk6O0vABzfZqgvMVN__eHeGy7eMPUFjuap8Am2dMUzZqCWHGZispnDqyFE6bMr0kFfBqzwuVJiyBJ4V6NvuS6Bfpt7-AXiCidV8yHBhep3dsDhjOIxsOTJ1fSjzGyic",
         "q":"0Ep6MLyDKyI2FU_jmCCAA9ri5th8oiM_3F9oZMu9sLWsZnKd3ubo9G7I4njIPUOsmdpu8Py_-wRHOdVAiYR-DPtcFWFuCn1s-LavrrcbcUWVU1cVPbXtNaxS6Rmg0xMA8_BX_Ed-j3wT6EnD9heFa1LlPGCn-lqOTjJQhbpzveE",
         "dp":"00DPFYcDXqwmt95LgGnuCgr67XIpR-pRwLFuTQw9yrO09SvVoN_uYgsUCrzWaadaCVVdjew8C0GCUYrvkufHmW7JQXVzskE1ztYrm1Phz46_66SnWL2wq-hbMI4RAodGwqvYFjHmKr7RfYc_8sFZtEnO-ig3cpVUaSbDSt9kp9k",
         "dq":"ujY7MyL1KTS2oSsAz9oOWGPxgmt8UP-ucfehvqse9MIWVKiXVtWc4hnA4icU7SB5SDquACgWAfV7L8rknYyjnDuDwWGPHTWwnFHGki4FDKkmrJEn3jmszdU3ckmFK-7LZUozfsjz7pcAvSRlWrcAgkhjxCytP_aBTotQzZ05KEE",
         "qi":"b8wAeuHREuWPeT1Uid1AiEBB7HqcInEH4HoSgEYA5kh6m-Z9sqYW0LFT-muJIC1EO8hyIz-MMwcahwjfaKz2-BRCc6xeLTLk0ivYU9S_0ciPPhw7ZUkGYuTmxBAzUTIkQ3sd-2d7Abqk_4s42E0_Z_RP2uG7WDDtXC5Wtivpdeo"
      },
      {
         "kty":"RSA",
         "kid":"0da91d15-f4aa-42fd-a35b-62a047c1f31e",
         "use":"sig",
         "alg":"RS256",
         "n":"w5INamZaWL0ZFBuu8xYVcN9zEey4191t_s0-rMnO9IpNFbC6V1QBt0BUCYuM0m2JpTLI8UMiCpGDHx4GIWuwXCUVVpuLgOxf4a2ewI7QmaVeKv4NMSUKdkyaZDlUpw5lpz_lyajmd4wW3WKnLFCoFq0cxfkNRdmz0wgcP9byk_YCkX_LKA_9bNBgrxOuGyLs3zcU8Mt-78iCO0_hycfc5nEfj2YSv7Dp-ohRiWKxy7hKwvB7PXqxstCDKq0weSDBDSCaGamfWTaVWkNRrM88K3xIXPLVTVxaVvgq9Avb20CACqFz-v1vqA3xwyck_uqQx_ynL8q4tTmbEqR7iE7E2Q",
         "e":"AQAB",
         "d":"MvKsfgh2BuIUU3G2zOr24PXFM2k46jtsOVHcvwS_3mLdHAZzNObUn5mpiucI45TXo-qsAHYduyUXRJb8v94fDpI2kd5ppEdv3wns7nsGCtDSzF5sr13X8OOZ4Pwyema8wqpZwYQ3rfMXzcqyhE_qyiWE9mogNA830oUtXtAvV63tdNuroPcMpUb1k76VzF7_mYzVXg6W0ggOYWd4Ekl1BaZGjle13TTv7MauVQlwdmyQywudiUxtZtaf2kxsk3yszuriKouEp8I4WM-2QRQ8vbw-sot66-Z2A3ry1hMDTRLT4OpkcmC8OpN6_5cNPKVVZuiriBgK6k9uf3isg3IsZQ",
         "p":"5XkZSUHFWoffBGy_YRbzS2uBfoqHi_lZRU3azcshDzOhMzgjE7RGXYiSRIveXnxusvs1cpNZaR6SFTnB2PUPGw8rF8qKsvwN9JlbSxUW6Y8KILoqMwIqFbKQkP2At4JoAKksEbWufX_WnQ4nE72s8oZWrhXnYdvQIijs2AzEf0s",
         "q":"2i2pHtJiPk6GSYzsl1gsRbkJf2J49MwyccLUleUrCOjwCcfZ1lIptfBmbrLCEKWBS0EiaWjzIpgcrn5IXoUGWmlDie9PoyB-OsaiIr2HUtYhNVgMfeWDa1f-XaAHI-0IpTsCPHVNhwCPaiIN1btBwYUMF_K-c-sOJOK2XUMH4es",
         "dp":"L5IFLeyWjwfvsakhm2z4jsAAnkz1gIz2dqmHHHZpqg8X8dhHXURX-fff6pncRVYiVLRDBjeJp4MQ4ZmRl_plYUSpuEriqewasIRCKrR6hXyDqvks2loug0T7NzN2RZShHtHzMtCpFZc01GYkr7D4c2Zp_bjIuL6qzQoS3072RTU",
         "dq":"x154L3kSW1tILQfA3t-svR3ERwpF-3RpGDlwJ3VdqOurBVUBg25bPS1rfPaOjcWfa5QejPJhUxhrBSzSlsS6NH4CQJZVUHyNvMnsORC2CwUvHV7TM2w6dinXf9iPDc45Wzub5IFQke_6HYL233sZMGySy4N_c8-0ghJFkN2C43c",
         "qi":"uzlkTTf0qYxlhGlqdF-JsPpvVaVEITuyIfNUl0lTOI8JTU7p-WkzPxy8d9kZVxCaiM_Of8C8JNvZ_WFDmbGnW8czCBqDpM3AykdmCEAT28zAAYgIan9ODhpOCzYA86M5G2bGEj-fNsxosIyupF5F1Nb1DXks6I3hpYVmOnuvuPA"
      },
      {
         "kty":"RSA",
         "kid":"e831a90d-d6bf-4a42-8c47-83d3ee5ab7ce",
         "use":"enc",
         "alg":"RSA1_5",
         "n":"o1OWo7p90tyRIOVcDFZREWMDKF1aSZM2aLE9yozwdRuty16w5-G-GfxEbK79D3z4LWkKf_4jy73ZWLXBcg_z4jUmj-MhHJTjgMWO8ovhOKm87UGqt81eRVJHNs10wmiq31XRYCfHbC70CnVmCByVR3Y4o6BjQzhjtgk1X6nM-I-4b5qHdmNoHYJFDLcHcGAjHv6azR2Ml7sF2nbAxIsSque_UN132EscxioMgabx8Nw4GdhzRMOt8d_Giq66aFKa7XH4aEhCcYnqxfYghRPZ4eLMfUFKO27ahvjfAtySTazDpLixPjI-xOkbSJxArGDFWf9deoE4rR-iSOi9Njl-5w",
         "e":"AQAB",
         "d":"aVtufUe8Ct3FsNbDviGgQfsA-nTd3UBdhMNw_MxWAPVN8zauH6b7nn-hFAr9q5QN9B0ibNZf-PoKzrLQiufHh1CKA5-cqHdTGqpWQBgDvS2hdds6aT7NZSJiVfPLMe2a7F7LpZ-DgUH4oxaALxNhKKwWdVbtMVua1r6x83fxhRUDlcEIRzyAKrQaUMCRMJMy7VVZ_oEo_bLyajstg6zbtqf-nBS1BUmAMQRVU366oWGSrWBPEM2ET26Tt9Rii7BHTv2ypo1owaPKGDvNDcYGD6bwd5EFAO1TNZaDQ0mJP20bK7c8xyNaPfnNDh90Vomv3me7mwvMA0lVABSiLAoaoQ",
         "p":"-zW2FnJhgxN0g8vThMCrwtLcOyANhjbAyvBOdyppqY3d70x2G4_m6f22O_oa3Ey5k-4THICckrfMlVk2TBC7rhcLUpjtBJ418LD73F2V2eWziv3WYTi9KJbS2NsufXdZGEqo7d-Qh0_P8Ri_xFraaIMKfpXhAhr2FAbS_W9kcUU",
         "q":"pnDfPuYviDnqEe5Y_6O8lyHdloPXnY3JwvxKMxUv3-yYtqHfqdQvKv4ZuDY7_pTTb4sdKhgy5Ud58jO48YbgqY3HKVth1YIPhxnb4mK4U7WrPxWoBK8pz1_Guvvh1xDbS6u5C9OrqdY0AiYYyG95lUzSdqyUi0F-4yWdTgbyFDs",
         "dp":"QtZ2nSBPu41Imex4WcDdsldiC0Uq9APLZfNsHR6mwFsjqpDAd_LgsG81tl2EGgs78RUN9q5tekf24eG7pZ9qIBa3h4FyxqDFn0WnrWkk_rW0AI4rJPDwu0Tt0o72nqFLjkAHFEtAbBAbNn2sQDUgGWCMQUPleybrREbQimfB5LU",
         "dq":"BK456PfqMEeIqJZuVEoTfKCMLbZpctnQ6bXUlFktLnvl04T72DfKV8griv2jdEZVJ9berBdgHwiCimgf9FLZsIr3JdeXCb0NmLwGbfhevKPoO-7s-ay_XUCRQyLgN_8WW6tpmcaLFkyay9Csc76GyccOSB4UU1I1MkgVg2M4nY8",
         "qi":"Y9D1SatYbCDAAMBxKs3JZhEpbqRKR87MQe4e0OKsJSHrrnVpsoEF1659HtRkm_rPaEsotKhic_KqpFmWg-deKKoaJiyDwtIT3mNBQRcEg-cizIddDyuC2MdmXqmLwAsjRAaEFtSgzbmpBluzJT3Rc7ruPpBrEbojZL698WbavWY"
      }
   ]
}
`
)

func TestJwtTokenStrategy_New(t *testing.T) {
	logger := log.NewNopLogger()

	var jwks *jose.JSONWebKeySet
	{
		jwks = &jose.JSONWebKeySet{}
		require.Nil(t, json.NewDecoder(strings.NewReader(testJwks)).Decode(jwks))
	}

	var strategy TokenStrategy
	{
		strategy = NewJwtTokenStrategy(jwks, jose.RS256, "test", 3600, 300, logger)
		require.NotNil(t, strategy)
	}

	var (
		err error
		tok string
	)
	{
		tok, err = strategy.New(context.Background(), &exported.Session{
			RequestId:   "AE7E43B6-E727-463A-B657-03163A17F411",
			ClientId:    "bar",
			Subject:     "foo",
			RedirectUri: "http://test.org/callback",
			AccessClaims: map[string]interface{}{
				"A": "a",
			},
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, tok)
	}
}

func TestJwtTokenStrategy_DeTokenize(t *testing.T) {
	logger := log.NewNopLogger()

	var jwks *jose.JSONWebKeySet
	{
		jwks = &jose.JSONWebKeySet{}
		require.Nil(t, json.NewDecoder(strings.NewReader(testJwks)).Decode(jwks))
	}

	var strategy TokenStrategy
	{
		strategy = NewJwtTokenStrategy(jwks, jose.RS256, "test", 3600, 300, logger)
		require.NotNil(t, strategy)
	}

	var tok string
	{
		tok, _ = strategy.New(context.Background(), &exported.Session{
			RequestId:   "AE7E43B6-E727-463A-B657-03163A17F411",
			ClientId:    "bar",
			Subject:     "foo",
			RedirectUri: "http://test.org/callback",
			AccessClaims: map[string]interface{}{
				"A": "a",
			},
		})
		require.NotEmpty(t, tok)
	}

	var (
		sess *exported.Session
		err  error
	)
	{
		sess, err = strategy.DeTokenize(context.Background(), tok)
		assert.Nil(t, err)
		assert.NotNil(t, sess)
	}

	{
		assert.Equal(t, "AE7E43B6-E727-463A-B657-03163A17F411", sess.RequestId)
		assert.Equal(t, "bar", sess.ClientId)
		assert.Equal(t, "foo", sess.Subject)
	}
}
