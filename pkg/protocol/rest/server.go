package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"github.com/dhanusaputra/somewhat-server/pkg/logger"
	"github.com/dhanusaputra/somewhat-server/pkg/protocol/rest/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := v1.RegisterSomewhatHandlerFromEndpoint(ctx, gwmux, "localhost:"+grpcPort, opts); err != nil {
		logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
	}
	mux.Handle("/", gwmux)

	fs := http.FileServer(http.Dir("./third_party/swagger-ui"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))

	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr: ":" + httpPort,
		// add handler with middleware
		Handler: cors.Default().Handler(
			middleware.AddRequestID(
				middleware.AddLogger(logger.Log, mux))),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	logger.Log.Info("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
