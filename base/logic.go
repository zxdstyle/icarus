package base

import (
	"context"
	"github.com/zxdstyle/icarus/policies"
	"github.com/zxdstyle/icarus/server/requests"
	"github.com/zxdstyle/icarus/server/responses"
)

type (
	logicModel interface {
		GetID() uint
	}

	Logic[M logicModel] interface {
		List(ctx context.Context, req requests.Request, response *responses.ApiResponse) error
		Show(ctx context.Context, primaryKey uint, req requests.Request, mo *M) error
		Create(ctx context.Context, mo *M) error
		Update(ctx context.Context, primaryKey uint, mo *M) error
		Destroy(ctx context.Context, primaryKey uint) error
	}

	BaseLogic[V logicModel] struct {
		repo Repository[V]
	}
)

var _ Logic[RepositoryModel] = &BaseLogic[RepositoryModel]{}

func NewLogic[V logicModel](repo Repository[V]) *BaseLogic[V] {
	return &BaseLogic[V]{repo: repo}
}

func (b BaseLogic[V]) List(ctx context.Context, request requests.Request, paginator *responses.ApiResponse) error {
	return b.repo.List(ctx, request, paginator)
}

func (b BaseLogic[V]) Show(ctx context.Context, id uint, req requests.Request, mo *V) error {
	return b.repo.Show(ctx, id, req, mo)
}

func (b BaseLogic[V]) Create(ctx context.Context, mo *V) error {
	return b.repo.Create(ctx, mo)
}

func (b BaseLogic[V]) Update(ctx context.Context, primaryKey uint, mo *V) error {
	return b.repo.Update(ctx, primaryKey, mo)
}

func (b BaseLogic[V]) Destroy(ctx context.Context, primaryKey uint) error {
	return b.repo.Destroy(ctx, primaryKey)
}

func (b BaseLogic[V]) Authorize(ctx context.Context, pos ...policies.Policy) error {
	for _, po := range pos {
		if err := po.Authorize(ctx); err != nil {
			return err
		}
	}
	return nil
}
