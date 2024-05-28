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

type Goals interface {
	Get() ([]*models.Goal, error)
	GetById(id int) (*models.Goal, error)
	Create(goal models.Goal) error
	Update(goal models.Goal) error
	Delete(id int) error
}

type Repository struct {
	Records
	Goals
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Records: NewRecordRepository(db),
		Goals:   NewGoalRepository(db),
	}
}
