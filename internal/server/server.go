package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-facade/internal/config"
	"github.com/ozonmp/bss-office-facade/internal/handlers"
	"github.com/ozonmp/bss-office-facade/internal/kafka"
	"github.com/ozonmp/bss-office-facade/internal/logger"
	"github.com/ozonmp/bss-office-facade/internal/model"
	"github.com/ozonmp/bss-office-facade/internal/repo"
	pb "github.com/ozonmp/bss-office-facade/pkg/bss-office-facade"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

// ConsumerServer сервер для сервиса-консюмера
type ConsumerServer struct {
	db *sqlx.DB
}

// NewConsumerServer создаёт новый сервер
func NewConsumerServer(db *sqlx.DB) *ConsumerServer {
	return &ConsumerServer{
		db: db,
	}
}

// Start запускает сервер
func (s *ConsumerServer) Start(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	metricsAddr := fmt.Sprintf("%s:%v", cfg.Metrics.Host, cfg.Metrics.Port)

	metricsServer := createMetricsServer(cfg)

	go func() {
		logger.InfoKV(ctx, "Metrics server is running", "addr", metricsAddr)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, "Failed running metrics server", "err", err)
			cancel()
		}
	}()

	isReady := &atomic.Value{}
	isReady.Store(false)

	statusServer := createStatusServer(cfg, isReady)

	go func() {
		statusAdrr := fmt.Sprintf("%s:%v", cfg.Status.Host, cfg.Status.Port)
		logger.InfoKV(ctx, "Status server is running", "addr", statusAdrr)
		if err := statusServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, "Failed running status server", "err", err)
		}
	}()

	officeRepo := repo.NewOfficeRepo(s.db)
	handler := handlers.NewEventHandler(officeRepo)

	err := kafka.StartConsuming(ctx, cfg.Kafka.Brokers, cfg.Kafka.Topic, cfg.Kafka.GroupID, handler.Handle)

	if err != nil {
		logger.FatalKV(ctx, "Failed starter consumer", "err", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.InfoKV(ctx, "signal.Notify", "sig", v)
	case done := <-ctx.Done():
		logger.InfoKV(ctx, "ctx.Done", "sig", done)
	}

	isReady.Store(false)

	if err := statusServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, "statusServer.Shutdown", "err", err)
	} else {
		logger.InfoKV(ctx, "statusServer shut down correctly")
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, "metricsServer.Shutdown", "err", err)
	} else {
		logger.InfoKV(ctx, "metricsServer shut down correctly")
	}

	return nil
}

func handleEvent(ctx context.Context, message *sarama.ConsumerMessage) error {
	var pbEvent pb.OfficeEvent
	err := proto.Unmarshal(message.Value, &pbEvent)

	if err != nil {
		return err
	}

	fmt.Printf("%#+v\n", model.ConvertPbToBssOfficeEvent(&pbEvent))

	return nil
}
