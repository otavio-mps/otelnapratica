// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"encoding/json"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/model"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
)

// SubscriptionHandler is an HTTP handler that performs CRUD operations for model.Subscription using a store.Subscription
type SubscriptionHandler struct {
	store         store.Subscription
	usersEndpoint string
	plansEndpoint string
}

// NewSubscriptionHandler returns a new SubscriptionHandler
func NewSubscriptionHandler(store store.Subscription, usersEndpoint string, plansEndpoint string) *SubscriptionHandler {
	return &SubscriptionHandler{
		store:         store,
		usersEndpoint: usersEndpoint,
		plansEndpoint: plansEndpoint,
	}
}

func (h *SubscriptionHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	subscriptions, err := h.store.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(subscriptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SubscriptionHandler) Create(w http.ResponseWriter, r *http.Request) {
	subscription := &model.Subscription{}
	if err := json.NewDecoder(r.Body).Decode(subscription); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// verify the user exists
	{
		user, _ := http.Get(h.usersEndpoint + "/" + subscription.UserID)
		if user.StatusCode != http.StatusOK {
			http.Error(w, "User not found", http.StatusBadRequest)
			return
		}
		defer user.Body.Close()
	}

	// verify the plan exists
	{
		plan, _ := http.Get(h.plansEndpoint + "/" + subscription.PlanID)
		if plan.StatusCode != http.StatusOK {
			http.Error(w, "Plan not found", http.StatusBadRequest)
			return
		}
		defer plan.Body.Close()
	}

	created, err := h.store.Create(r.Context(), subscription)
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

func (h *SubscriptionHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	subscription, err := h.store.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if subscription == nil {
		http.Error(w, "Subscription not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(subscription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SubscriptionHandler) Update(w http.ResponseWriter, r *http.Request) {
	subscription := &model.Subscription{}
	if err := json.NewDecoder(r.Body).Decode(subscription); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedSubscription, err := h.store.Update(r.Context(), subscription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(updatedSubscription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SubscriptionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.store.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
