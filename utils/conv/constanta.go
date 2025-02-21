package conv

import "errors"

type ContextKey string

const (
	CtxUserAgent = ContextKey("user-agent")
)

const (
	MessageSuccess = "Success!"
)

var (
	ErrInternalServerError  = errors.New("internal server error")
	ErrNotFound             = errors.New("not found")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
)
