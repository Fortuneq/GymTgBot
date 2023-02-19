package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gogrpc "google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	cfg *Config
}

func (a *App) Run() error {
	err := validator.New().Struct(a.cfg)
	if err != nil {
		return fmt.Errorf("valdiate config: %w", err)
	}

	atom := zap.NewAtomicLevel()
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		atom,
	)
	logger := zap.New(zapCore)
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			return
		}
	}(logger)

	log := logger.Sugar()
	atom.SetLevel(zapcore.Level(*a.cfg.Logger.Level))
	log.Infof("logger initialized successfully")

	grpcServer := gogrpc.NewServer()
	defer grpcServer.GracefulStop()

	go func() {
		lis, err := net.Listen("tcp", a.cfg.Server.Host+a.cfg.Server.Port)
		if err != nil {
			log.Fatalf("tcp sock: %s", err.Error())
		}
		defer func(lis net.Listener) {
			err = lis.Close()
			if err != nil {

			}
		}(lis)

		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("GRPC server: %s", err.Error())
		}
	}()

	app := fiber.New()

	app.Use(helmet.New())
	app.Use(recover2.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	apiGroup := app.Group("api")

	apiGroup.Get("register", )

	log.Debug("Started GRPC server")

	log.Debug("Application has started")

	exit := make(chan os.Signal, 2)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	log.Info("Application has been shut down")

	return nil

}

func New(cfg *Config) *App {
	return &App{cfg}
}
