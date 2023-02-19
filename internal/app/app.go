package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gym-bot/internal/controller"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	tele "gopkg.in/telebot.v3"

	"gym-bot/pkg/postgres"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	psqlDB, err := postgres.InitPsqlDB(a.cfg)
	if err != nil {
		log.Fatalf("PostgreSQL init error: %s", err)
	} else {
		log.Infof("PostgreSQL connected, status: %#v", psqlDB.Stats())
	}
	defer func(psqlDB *sqlx.DB) {
		err = psqlDB.Close()
		if err != nil {
			log.Infof(err.Error())
		} else {
			log.Info("PostgreSQL closed properly")
		}
	}(psqlDB)

	prefBot := tele.Settings{
		Token:  a.cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	Bot, err := tele.NewBot(prefBot)
	if err != nil {
		log.Fatal(err)
	}

	gymController := controller.NewBotController(Bot)
	gymController.Start()

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
