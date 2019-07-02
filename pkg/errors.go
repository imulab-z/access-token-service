package pkg

import "fmt"

type ServiceError struct {
	Err         string            `json:"error"`
	Description string            `json:"error_description"`
	StatusCode  int               `json:"-"`
	Headers     map[string]string `json:"-"`
}

func (e ServiceError) Error() string {
	return e.Err
}

func (e ServiceError) IsInvalidGrantError() bool {
	return e.Err == "invalid_grant"
}

func ErrTokenExpired(token string) error {
	return &ServiceError{
		Err:         "invalid_grant",
		Description: fmt.Sprintf("token [%s] has expired.", token),
		StatusCode:  400,
	}
}

func ErrInvalidToken(token string) error {
	return &ServiceError{
		Err:         "invalid_grant",
		Description: fmt.Sprintf("token [%s] is invalid.", token),
		StatusCode:  400,
	}
}

func ErrInvalidRequest(detail string) error {
	return &ServiceError{
		Err:         "invalid_request",
		Description: detail,
		StatusCode:  400,
	}
}

func ErrServer(reason string) error {
	return &ServiceError{
		Err:         "server_error",
		Description: reason,
		StatusCode:  500,
	}
}

