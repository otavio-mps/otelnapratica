// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package model

import "time"

type Payment struct {
	ID             string    `json:"id"`
	SubscriptionID string    `json:"subscription_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Version        int64     `json:"version"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
