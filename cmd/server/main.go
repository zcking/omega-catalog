package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	omega_catalogv1 "github.com/zcking/omega-catalog/gen/go/omega_catalog/v1"
	"github.com/zcking/omega-catalog/internal"
	"github.com/zcking/omega-catalog/internal/database"
)

var (
	// configFile is the path to the configuration file
	configFile = flag.String("config", "conf/server.yaml", "path to the server configuration file")

	// logger is a global zap logger. It is initialized once in the initLogger
	// function then zap.L() can be used to retrieve it anywhere in the code.
	logger       *zap.Logger
	loggerConfig zap.Config

	// tracer is a global OpenTelemetry tracer.
	tracer = otel.Tracer("omega-catalog")
)

func initLogger() {
	var err error
	// Configure the logger to use a production configuration
	loggerConfig = zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err = loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}

func main() {
	// Parse CLI flags
	flag.Parse()

	// Initialize zap logger
	initLogger()
	defer logger.Sync()

	// Load the configuration file
	config, err := internal.NewConfigFromPath(*configFile)
	if err != nil {
		logger.Fatal("failed to load configuration file", zap.Error(err))
	}

	// Initialize connection to the database
	dsn := database.NewDSN(
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Properties,
	)
	logger.Info("connecting to database", zap.String("dsn", dsn)) // TODO: remove this line
	db, err := database.NewDatabase(dsn)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}

	grpcAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	restAddr := fmt.Sprintf("%s:%d", config.RestGateway.Host, config.RestGateway.Port)

	// Create a TCP listener for the gRPC server
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Fatal("failed to create (gRPC) listener", zap.Error(err))
	}

	// Create a gRPC server and attach our implementation
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
		)),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}
	grpcServer := grpc.NewServer(opts...)
	impl, err := internal.NewImpl(db)
	if err != nil {
		logger.Fatal("failed to create server", zap.Error(err))
	}
	omega_catalogv1.RegisterOmegaCatalogServiceServer(grpcServer, impl)

	// Catch interrupt signal to gracefully shutdown the server
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Set up OpenTelemetry.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	otelShutdown, err := internal.SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	go func() {
		sig := <-signalChan
		logger.Info(
			"received signal. shutting down gRPC server...",
			zap.String("signal", sig.String()),
		)
		grpcServer.GracefulStop()
		cancel() // make sure the context is stopped too
		if err := impl.Close(); err != nil {
			logger.Fatal("failed to properly close server", zap.Error(err))
		}
	}()

	// Serve the gRPC server, in a separate goroutine to avoid blocking
	go func() {
		logger.Info("gRPC server listening on " + grpcAddr)
		logger.Fatal("server stopped.", zap.Error(grpcServer.Serve(lis)))
	}()

	// Now setup the gRPC Gateway, a REST proxy to the gRPC server
	conn, err := grpc.NewClient(
		grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		logger.Fatal("failed to create gRPC client", zap.Error(err))
	}

	mux := runtime.NewServeMux()

	// Expose the zap configuration for getting or changing the log level at runtime
	mux.HandlePath(http.MethodGet, "/log/level", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		loggerConfig.Level.ServeHTTP(w, r)
	})
	mux.HandlePath(http.MethodPut, "/log/level", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		loggerConfig.Level.ServeHTTP(w, r)
	})

	err = omega_catalogv1.RegisterOmegaCatalogServiceHandler(context.Background(), mux, conn)
	if err != nil {
		logger.Fatal("failed to register gRPC gateway", zap.Error(err))
	}

	// Start HTTP server to proxy requests to gRPC server
	gwServer := &http.Server{
		Addr:    restAddr,
		Handler: mux,
	}
	logger.Info("gRPC Gateway listening", zap.String("address", restAddr))
	logger.Fatal("gRPC Gateway stopped.", zap.Error(gwServer.ListenAndServe()))
}
