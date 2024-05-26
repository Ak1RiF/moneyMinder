package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func NewPostgresDB(cfg Config) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DB, cfg.SSLMode))

	if err != nil {
		return nil, err
	}
	return db, nil
}
