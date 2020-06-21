package apifrontend

import (
	"context"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/model"
)

type Manager interface {
	CreateAPIFrontend(ctx context.Context, schema int, path string, method string,
		param []byte, body []byte) (uint32, error)

	GetAPIFrontend(ctx context.Context, id uint32) (model.APIFrontend, error)

	UpdateFrontend(ctx context.Context, id uint32, schema int, path string, method string,
		param []byte, body []byte)
}
