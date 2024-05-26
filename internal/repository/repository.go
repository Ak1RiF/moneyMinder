package repository

import (
	"MoneyMinder/internal/models"
	"github.com/jackc/pgx/v5"
)

type Records interface {
	GetByType(typeStr string) ([]*models.Record, error)
	GetById(id int) (*models.Record, error)
	Create(entity models.Record) error
	Update(id int, entity models.Record) error
}

type Repository struct {
	Records
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Records: NewRecordRepository(db),
	}
}
