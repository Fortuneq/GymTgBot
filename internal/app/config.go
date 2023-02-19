package app

import "time"

type (
	Config struct {
		Server   Server   `yaml:"Server" validate:"required"`
		Logger   Logger   `yaml:"Logger" validate:"required"`
		Postgres Postgres `yaml:"Postgres" validate:"required"`
		Telegram Telegram `yaml:"Telegram" validate:"required"`
	}
	Telegram struct {
		Token string `yaml:"token" validate:"required"`
	}

	Server struct {
		Host string `yaml:"Host" validate:"required"`
		Port string `yaml:"Port" validate:"required"`
	}
	Postgres struct {
		Host            string `validate:"required"`
		Port            string `validate:"required"`
		User            string `validate:"required"`
		Password        string `validate:"required"`
		DBName          string `validate:"required"`
		SSLMode         string `validate:"required"`
		ApplicationName string `validate:"required"`
		PGDriver        string `validate:"required"`
		Settings        struct {
			MaxOpenConns    int           `validate:"required,min=1"`
			ConnMaxLifetime time.Duration `validate:"required,min=1"`
			MaxIdleConns    int           `validate:"required,min=1"`
			ConnMaxIdleTime time.Duration `validate:"required,min=1"`
		}
	}

	Logger struct {
		Level *int8 `yaml:"Level" validate:"required"`
	}
)
