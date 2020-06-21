package api

import (
	"context"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/model"
)

type Manager interface {
	CreateAPI(ctx context.Context, name string, domain string, authType int, qps int,
		groupName string, groupID uint32, frontendID uint32, backendID uint32, resultID uint32) (model.API, error)
}
