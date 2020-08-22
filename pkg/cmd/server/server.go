package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	pb "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"github.com/dhanusaputra/somewhat-server/pkg/logger"
	"github.com/dhanusaputra/somewhat-server/pkg/protocol/grpc"
	"github.com/dhanusaputra/somewhat-server/pkg/protocol/rest"
	v1 "github.com/dhanusaputra/somewhat-server/pkg/service/v1"
	"github.com/dhanusaputra/somewhat-server/util/jsonutil"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9090", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "8080", "HTTP port to bind")
	flag.IntVar(&cfg.LogLevel, "log-level", -1, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "2006-01-02T15:04:05.999999999Z07:00", "Print time format for logger e.g. 006-01-02T15:04:05Z07:00")
	flag.Parse()

	port := os.Getenv("PORT")

	if port == "" {
		port = cfg.HTTPPort
	}

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(port) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", port)
	}

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	var data interface{}
	err := jsonutil.ReadFile("./api/json/v1/db.json", &data)
	if err != nil {
		return err
	}

	var userData []pb.User
	err = jsonutil.ReadFile("./api/json/v1/user.json", &userData)
	if err != nil {
		return err
	}

	v1API := v1.NewServer(data.(map[string]interface{}), userData)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, port)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
