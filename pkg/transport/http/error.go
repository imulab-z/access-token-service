package http

import (
	"context"
	"encoding/json"
	"github.com/imulab-z/access-token-service/pkg"
	"net/http"
)

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	var se *pkg.ServiceError
	{
		switch err.(type) {
		case *pkg.ServiceError:
			se = err.(*pkg.ServiceError)
		default:
			se = pkg.ErrServer(err.Error()).(*pkg.ServiceError)
		}
	}

	w.WriteHeader(se.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(se)
}

