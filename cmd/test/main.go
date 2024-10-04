package main

import (
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test/infrastructure/db"
	"test/infrastructure/logger"
	"test/internal/metrics"
	"test/internal/repositories"
	swagger "test/pkg/clients/garantex"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"test/api/test"
	"test/config"
	notifimpl "test/internal/app/test"
)

func main() {
	logger.InitLogger()
	metrics.Init()
	_ = godotenv.Load(".env")
	err := config.NewConf()
	if err != nil {
		zap.L().Fatal("failed to load config", zap.Error(err))
	}

	zap.L().Info("starting server")
	listener, err := net.Listen("tcp", ":"+config.Cnf.TcpPort)
	if err != nil {
		panic(err)
	}

	conn, err := db.New(&config.Cnf)
	if err != nil {
		zap.L().Fatal("app - Run - postgres.New", zap.Error(err))
	}
	defer conn.Close()

	tnq := repositories.NewRateQuery(*conn)

	garantexClient := swagger.NewAPIClient(swagger.NewConfiguration())

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	grpc_prometheus.Register(s)

	doneCh := make(chan struct{})
	test.RegisterTestServer(s, notifimpl.NewTest(tnq, garantexClient.DepthApi))
	go func() {
		if err = s.Serve(listener); err != nil {
			zap.L().Fatal("failed to serve", zap.Error(err))
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health",
		func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Cnf.HttpPort), nil))
	}()

	go gracefulShutdown(s, 5*time.Second, doneCh)

	<-doneCh
	zap.L().Info("Server stopped")
}

func gracefulShutdown(grpcServer *grpc.Server, timeout time.Duration, doneCh chan struct{}) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	zap.L().Info("Received shutdown signal")

	gracefulStop := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(gracefulStop)
	}()

	select {
	case <-gracefulStop:
		zap.L().Info("gRPC server shut down gracefully")
	case <-time.After(timeout):
		zap.L().Info("gRPC server did not shut down in time, forcing shutdown")
		grpcServer.Stop()
	}

	close(doneCh)
}
