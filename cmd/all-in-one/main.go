// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/telemetry"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
)

func main() {
	configFlag := flag.String("config", "", "path to the config file")
	flag.Parse()
	telemetry.InitTelemetry()

	c, _ := config.LoadConfig(*configFlag)

	mux := http.NewServeMux()

	ctx := context.Background()

	logExporter, err := otlploghttp.New(ctx, otlploghttp.WithInsecure())
	if err != nil {
		panic("failed to initialize exporter")
	}

	lp := log.NewLoggerProvider(
		log.WithProcessor(
			log.NewBatchProcessor(logExporter),
		),
	)

	defer lp.Shutdown(ctx)

	global.SetLoggerProvider(lp)

	logger := otelslog.NewLogger("teste")

	logger.Debug("Teste de log")

	// starts the gRPC server
	lis, _ := net.Listen("tcp", c.Server.Endpoint.GRPC)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	logger.Info("Server Up")

	{
		a := app.NewUser(&c.Users)
		a.RegisterRoutes(mux)
	}

	{
		a := app.NewPlan(&c.Plans)
		a.RegisterRoutes(mux, grpcServer)
	}

	{
		a, err := app.NewPayment(&c.Payments)
		if err != nil {
			panic(err)
		}
		a.RegisterRoutes(mux)
		defer func() {
			_ = a.Shutdown()
		}()
	}

	{
		a := app.NewSubscription(&c.Subscriptions)
		a.RegisterRoutes(mux)
	}

	go func() {
		_ = grpcServer.Serve(lis)
	}()

	_ = http.ListenAndServe(c.Server.Endpoint.HTTP, mux)

}
