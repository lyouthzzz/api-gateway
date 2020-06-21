package apiserver

import (
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/apiservice"
)

type server struct {
	host  string
	port  int
	debug bool
	svc   *apiservice.Service
}

func NewServer() *server {
	return &server{
		host:  "",
		port:  0,
		debug: false,
		svc:   nil,
	}
}
