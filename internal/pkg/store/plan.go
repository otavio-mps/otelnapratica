package store

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
)

type Plan interface {
	Get(ctx context.Context, id string) (*model.Plan, error)
	Create(ctx context.Context, plan *model.Plan) (*model.Plan, error)
	Update(ctx context.Context, plan *model.Plan) (*model.Plan, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.Plan, error)
}
