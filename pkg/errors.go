package pkg

import (
	"fmt"
	"github.com/imulab-z/common/errors"
)

func ErrTokenExpired() error {
	return errors.InvalidGrant("access token has expired")
}

func ErrInvalidToken() error {
	return errors.InvalidGrant("access token is invalid")
}

func ErrInvalidRequest(detail string) error {
	return errors.InvalidRequest(detail)
}

func ErrParameterRequired(param string) error {
	return errors.InvalidRequest(fmt.Sprintf("parameter <%s> is required", param))
}

func ErrServer(reason string) error {
	return errors.ServerError(reason)
}

