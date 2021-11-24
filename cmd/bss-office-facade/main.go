package main

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-office-facade/internal/logger"
	"github.com/ozonmp/bss-office-facade/internal/metrics"
	"log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozonmp/bss-office-facade/internal/config"
	"github.com/ozonmp/bss-office-facade/internal/database"
	"github.com/ozonmp/bss-office-facade/internal/server"
	"github.com/ozonmp/bss-office-facade/internal/tracer"
)

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatalf("Failed init configuration:%s", err)
	}

	cfg := config.GetConfigInstance()

	syncLogger := logger.InitLogger(ctx, &cfg)
	defer syncLogger()

	logger.InfoKV(ctx, fmt.Sprintf("Starting service: %s", cfg.Project.Name),
		"version", cfg.Project.Version,
		"commitHash", cfg.Project.CommitHash,
		"debug", cfg.Project.Debug,
		"environment", cfg.Project.Environment,
	)

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	metrics.InitMetrics(&cfg)

	db, err := database.NewPostgres(ctx, dsn, cfg.Database.Driver)
	if err != nil {
		logger.FatalKV(ctx, "Failed init postgres", "err", err)
	}
	defer db.Close()

	tracing, err := tracer.NewTracer(ctx, &cfg)
	if err != nil {
		logger.FatalKV(ctx, "Failed init tracing", "err", err)

		return
	}
	defer tracing.Close()

	if err = server.NewConsumerServer(db).Start(ctx, &cfg); err != nil {
		logger.FatalKV(ctx, "Failed creating consumer server", "err", err)

		return
	}
}
