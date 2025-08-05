package errors

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrUnimplemented   = errors.New("unimplemented")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrUnavailable     = errors.New("service unavailable")
	ErrUnknown         = errors.New("unknown")
	ErrInternal        = errors.New("internal")
	ErrUnsupported     = errors.New("unsupported")
	ErrTimeOut         = errors.New("request timeout")
)
