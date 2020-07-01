package auth

import "errors"

var (
	ErrAuthParam  = errors.New("auth: param error")
	ErrAuthVerify = errors.New("auth: verify error")
)
