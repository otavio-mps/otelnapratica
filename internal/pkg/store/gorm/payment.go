// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package gorm

import (
	"context"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"gorm.io/gorm"
)

type Payment struct {
	db *gorm.DB
}

func NewPaymentStore(db *gorm.DB) store.Payment {
	return &Payment{db: db}
}

func (p *Payment) Get(ctx context.Context, id string) (*model.Payment, error) {
	ret := &model.Payment{}
	_ = p.db.WithContext(ctx).Model(ret).First(&ret, "id = ?", id)
	return ret, nil
}

func (p *Payment) Create(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
	res := p.db.WithContext(ctx).Create(&payment)
	return payment, res.Error
}

func (p *Payment) Update(ctx context.Context, payment *model.Payment) (*model.Payment, error) {
	res := p.db.WithContext(ctx).Save(&payment)
	return payment, res.Error
}

func (p *Payment) Delete(ctx context.Context, id string) error {
	_ = p.db.WithContext(ctx).Delete(&model.Payment{}, "id = ?", id)
	return nil
}

func (p *Payment) List(ctx context.Context) ([]*model.Payment, error) {
	var ret []*model.Payment
	_ = p.db.WithContext(ctx).Find(&ret)
	return ret, nil
}
