// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	userhttp "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/handler/http"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/pkg/store/memory"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

)

type User struct {
	Handler *userhttp.UserHandler
	Store   store.User
}

func NewUser(*config.Users) *User {
	store := memory.NewUserStore()
	return &User{
		Handler: userhttp.NewUserHandler(store),
		Store:   store,
	}
}

func (a *User) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /users", otelhttp.NewHandler(http.HandlerFunc(a.Handler.List), "GET /users"))
	mux.Handle("POST /users", otelhttp.NewHandler(http.HandlerFunc(a.Handler.Create), "POST /users"))
	mux.Handle("GET /users/{id}", otelhttp.NewHandler(http.HandlerFunc(a.Handler.Get), "GET /users/{id}"))
	mux.Handle("PUT /users/{id}", otelhttp.NewHandler(http.HandlerFunc(a.Handler.Update), "PUT /users/{id}"))
	mux.Handle("DELETE /users/{id}", otelhttp.NewHandler(http.HandlerFunc(a.Handler.Delete), "DELETE /users/{id}"))
}
