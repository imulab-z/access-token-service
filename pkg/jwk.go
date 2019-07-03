package pkg

import (
	"gopkg.in/square/go-jose.v2"
)

func selectSignatureKey(keySet *jose.JSONWebKeySet, alg jose.SignatureAlgorithm, clientId string) *jose.JSONWebKey {
	candidates := make([]jose.JSONWebKey, 0)
	for _, k := range keySet.Keys {
		if k.Use == "sig" && k.Algorithm == string(alg) {
			candidates = append(candidates, k)
		}
	}

	switch len(candidates) {
	case 0:
		return nil
	case 1:
		return &candidates[0]
	default:
		if len(clientId) == 0 {
			return &candidates[0]
		}

		var sum int
		{
			sum = 0
			for _, c := range []rune(clientId) {
				sum += int(c)
			}
		}

		return &candidates[sum % len(candidates)]
	}
}
