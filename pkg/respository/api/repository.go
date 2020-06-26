package api

import (
	"context"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/model"
)

type IterateApiCallBack func(api model.API) error

type Repository interface {
	GetApi(ctx context.Context, filter *FilterParam) (model.API, error)
	CreateApi(ctx context.Context, meta *MetaParam) (model.API, error)
	UpdateApi(ctx context.Context, id uint32, meta *MetaParam) error
	DeleteApi(ctx context.Context, id uint32) error
	IterateApi(ctx context.Context, filter *FilterParam, cb IterateApiCallBack) error
}
