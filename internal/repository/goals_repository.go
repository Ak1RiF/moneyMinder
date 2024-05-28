package repository

import (
	"MoneyMinder/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
)

type GoalRepository struct {
	db *pgx.Conn
}

func NewGoalRepository(db *pgx.Conn) *GoalRepository {
	return &GoalRepository{db: db}
}

func (r *GoalRepository) Get() ([]*models.Goal, error) {
	var goals []*models.Goal
	query := `SELECT id, title, amount, total_contributed, date_create, date_completion FROM goals`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var goal models.Goal

		if err := rows.Scan(&goal.Id, &goal.Title, &goal.Amount, &goal.TotalContributed, &goal.DateCreate, &goal.DateCompletion); err != nil {
			continue
		}
		goals = append(goals, &goal)
	}

	return goals, nil
}

func (r *GoalRepository) GetById(id int) (*models.Goal, error) {
	var goal models.Goal
	query := `SELECT id, title, amount, total_contributed, date_create, date_completion FROM goals WHERE id=$1`

	if err := r.db.QueryRow(context.Background(), query, id).Scan(&goal.Id, &goal.Title, &goal.Amount, &goal.TotalContributed, &goal.DateCreate, &goal.DateCompletion); err != nil {
		return nil, err
	}
	return &goal, nil
}

func (r *GoalRepository) Create(goal models.Goal) error {
	query := `INSERT INTO goals (title, amount, date_create, date_completion) VALUES ($1,$2,$3,$4)`
	if _, err := r.db.Exec(context.Background(), query, goal.Title, goal.Amount, goal.DateCreate, goal.DateCompletion); err != nil {
		return err
	}
	return nil
}

func (r *GoalRepository) Update(goal models.Goal) error {
	query := `UPDATE goals SET title=$1, amount=$2, date_completion=$3 WHERE id=$4`
	if _, err := r.db.Exec(context.Background(), query, goal.Title, goal.Amount, goal.DateCompletion, goal.Id); err != nil {
		return err
	}
	return nil
}
func (r *GoalRepository) Delete(id int) error {
	query := `DELETE FROM goals WHERE id=$1`
	if _, err := r.db.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}
