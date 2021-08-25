//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	polvo_client "pkg.aiocean.dev/getpolvo/internal/polvo-client"
	"pkg.aiocean.dev/getpolvo/internal/server"
	"pkg.aiocean.dev/serviceutil/logger"
)

func InitializeHandler(ctx context.Context) (*server.Server, error) {
	wire.Build(
		polvo_client.GetPolvoClient,
		logger.NewLogger,
		server.WireSet,
	)

	return nil, nil
}
