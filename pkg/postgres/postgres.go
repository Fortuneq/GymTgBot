package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib" // driver
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	ConnString string `yaml:"conn_string" validate:"required"`
	// MaxOpenConns    int           `yaml:"max_open_conns" validate:"required"`
	// ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" validate:"required"`
	// MaxIdleConns    int           `yaml:"max_idle_conns" validate:"required"`
	// ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" validate:"required"`
}

func New(connString string, cfg Postgres) (db *sqlx.DB, err error) {
	db, err = sqlx.Connect("pgx", connString)
	if err != nil {
		return nil, err
	}

	// db.SetMaxOpenConns(cfg.MaxOpenConns)
	// db.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Second)
	// db.SetMaxIdleConns(cfg.MaxIdleConns)
	// db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime * time.Second)

	return db, nil
}
