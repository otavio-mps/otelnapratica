// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
)

type Payment interface {
	Get(ctx context.Context, id string) (*model.Payment, error)
	Create(ctx context.Context, user *model.Payment) (*model.Payment, error)
	Update(ctx context.Context, user *model.Payment) (*model.Payment, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.Payment, error)
}
