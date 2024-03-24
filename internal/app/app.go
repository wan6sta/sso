package app

import (
	grpcapp "github.com/wan6sta/sso/internal/app/grpc"
	"github.com/wan6sta/sso/internal/services/auth"
	sqlite "github.com/wan6sta/sso/internal/storage/sqllite"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
