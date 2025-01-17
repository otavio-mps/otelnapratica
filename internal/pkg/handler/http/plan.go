// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"encoding/json"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
)

// PlanHandler is an HTTP handler that performs CRUD operations for model.Plan using a store.Plan
type PlanHandler struct {
	store store.Plan
}

// NewPlanHandler returns a new PlanHandler
func NewPlanHandler(store store.Plan) *PlanHandler {
	return &PlanHandler{
		store: store,
	}
}

func (h *PlanHandler) List(w http.ResponseWriter, r *http.Request) {
	plans, err := h.store.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(plans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PlanHandler) Create(w http.ResponseWriter, r *http.Request) {
	plan := &model.Plan{}
	if err := json.NewDecoder(r.Body).Decode(plan); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	created, err := h.store.Create(r.Context(), plan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(created)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PlanHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	plan, err := h.store.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(plan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PlanHandler) Update(w http.ResponseWriter, r *http.Request) {
	plan := &model.Plan{}
	if err := json.NewDecoder(r.Body).Decode(plan); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updated, err := h.store.Update(r.Context(), plan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PlanHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.store.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
