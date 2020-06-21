package apibackend

import (
	"context"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/model"
)

type Manager interface {
	CreateAPIBackend(ctx context.Context, schema int, host, path, method string, timeout int64,
		param []byte, body []byte) (uint32, error)

	GetAPIBackend(ctx context.Context, id uint32) (model.APIFrontend, error)

	UpdateBackend(ctx context.Context, id uint32, schema int, host, path, method string, timeout uint64,
		param, body []byte)
}
