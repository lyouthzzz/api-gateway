package auth

import (
	"net/http"
)

type Auth interface {
	Add(r *http.Request) error
	Verify(r *http.Request) error
}
