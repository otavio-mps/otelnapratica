// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
)

type User interface {
	Get(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.User, error)
}
