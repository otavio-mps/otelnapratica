// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
)

type Subscription interface {
	Get(ctx context.Context, id string) (*model.Subscription, error)
	Create(ctx context.Context, user *model.Subscription) (*model.Subscription, error)
	Update(ctx context.Context, user *model.Subscription) (*model.Subscription, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*model.Subscription, error)
}
