// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"pkg.aiocean.dev/getpolvo/internal/polvo-client"
	"pkg.aiocean.dev/getpolvo/internal/server"
	"pkg.aiocean.dev/serviceutil/logger"
)

// Injectors from wire.go:

func InitializeHandler(ctx context.Context) (*server.Server, error) {
	zapLogger, err := logger.NewLogger(ctx)
	if err != nil {
		return nil, err
	}
	polvoServiceClient, err := polvo_client.GetPolvoClient(ctx)
	if err != nil {
		return nil, err
	}
	serverServer := server.NewServer(zapLogger, polvoServiceClient)
	return serverServer, nil
}
