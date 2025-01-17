// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
)

type inMemorySubscription struct {
	store map[string]*model.Subscription
}

func NewSubscriptionStore() store.Subscription {
	return &inMemorySubscription{
		store: make(map[string]*model.Subscription),
	}
}

func (u *inMemorySubscription) Get(_ context.Context, id string) (*model.Subscription, error) {
	return u.store[id], nil
}

func (u *inMemorySubscription) Create(_ context.Context, user *model.Subscription) (*model.Subscription, error) {
	u.store[user.ID] = user
	return user, nil
}

func (u *inMemorySubscription) Update(_ context.Context, user *model.Subscription) (*model.Subscription, error) {
	u.store[user.ID] = user
	return user, nil
}

func (u *inMemorySubscription) Delete(_ context.Context, id string) error {
	delete(u.store, id)
	return nil
}

func (u *inMemorySubscription) List(_ context.Context) ([]*model.Subscription, error) {
	users := make([]*model.Subscription, 0, len(u.store))
	for _, user := range u.store {
		users = append(users, user)
	}
	return users, nil
}
