// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package model

import "time"

type Plan struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Price       int32     `json:"price"`
	Description string    `json:"description"`
	Version     int32     `json:"version"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
