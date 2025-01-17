// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/telemetry"
)

func main() {
	configFlag := flag.String("config", "", "path to the config file")
	flag.Parse()

	telemetry.InitTelemetry()
	c, _ := config.LoadConfig(*configFlag)

	a := app.NewUser(&c.Users)
	a.RegisterRoutes(http.DefaultServeMux)
	_ = http.ListenAndServe(c.Server.Endpoint.HTTP, http.DefaultServeMux)
}
