package apiserver

import (
	"fmt"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/apiservice"
)

type Conf struct {
	Host  string
	Port  int
	Svc   *apiservice.Service
	Debug bool
}

func NewServer(conf Conf) *server {
	return &server{
		host:  conf.Host,
		port:  conf.Port,
		debug: conf.Debug,
		svc:   conf.Svc,
	}
}

func (sv *server) Setup() error {
	engine := sv.BuildRouter()
	err := engine.Run(fmt.Sprintf("%v:%v", sv.host, sv.port))
	return err
}
