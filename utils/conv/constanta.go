package conv

import "errors"

type contextKey string

const (
	CtxUserAgent = contextKey("user-agent")
)

const (
	MessageSuccess = "Success!"
)

var (
	ErrInternalServerError  = errors.New("internal server error")
	ErrNotFound             = errors.New("data not found")
	ErrUserAlreadyExist     = errors.New("user already exist")
	ErrBadParamInput        = errors.New("given param is not valid")
	ErrWrongEmailOrPassword = errors.New("wrong email/password")
)
