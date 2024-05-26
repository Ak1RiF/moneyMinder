package repository

import (
	"MoneyMinder/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
)

type RecordsRepository struct {
	db *pgx.Conn
}

func NewRecordRepository(db *pgx.Conn) *RecordsRepository {
	return &RecordsRepository{db: db}
}

func (r *RecordsRepository) GetByType(typeStr string) ([]*models.Record, error) {
	var records []*models.Record
	query := `SELECT id, name, type, amount, description FROM records WHERE type = $1`

	rows, err := r.db.Query(context.Background(), query, typeStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record models.Record

		if err := rows.Scan(&record.Id, &record.Name, &record.Type, &record.Amount, &record.Description); err != nil {
			continue
		}

		records = append(records, &record)
	}

	return records, nil
}

func (r *RecordsRepository) GetById(id int) (*models.Record, error) {
	var record models.Record

	query := `SELECT id, name, type, amount, description FROM records WHERE id = $1`

	if err := r.db.QueryRow(context.Background(), query, id).Scan(&record.Id, &record.Name, &record.Type, &record.Amount, &record.Description); err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *RecordsRepository) Create(entity models.Record) error {
	query := `INSERT INTO records (name, type, amount, description) VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(context.Background(), query, entity.Name, entity.Type, entity.Amount, entity.Description)
	if err != nil {
		return err
	}

	return nil
}
func (r *RecordsRepository) Update(id int, entity models.Record) error {
	query := `UPDATE records SET name=$1, type=$2, amount=$3, description=$4 WHERE id=$5`

	_, err := r.db.Exec(context.Background(), query, entity.Name, entity.Type, entity.Amount, entity.Description, id)
	if err != nil {
		return err
	}

	return nil
}
