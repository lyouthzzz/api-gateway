package api

import "errors"

var (
	ErrApiNotFound       = errors.New("API: not found")
	EerrApiAlreadyExixts = errors.New("API: already exists")
)
