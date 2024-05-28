package service

import (
	"MoneyMinder/internal/dtos"
	"MoneyMinder/internal/repository"
)

type Records interface {
	GetProfitRecords() ([]*dtos.RecordOutput, error)
	GetExpenseRecords() ([]*dtos.RecordOutput, error)
	AddRecord(inputBody dtos.RecordInput) error
	UpdateRecord(recordId int, updateBody dtos.RecordInput) error
}

type Goals interface {
	GetAllGoals() ([]*dtos.GoalOutput, error)
	GetGoalById(id int) (*dtos.GoalOutput, error)
	CreateNewGoal(form dtos.CreateGoal) error
	UpdateFieldGoal(goalId int, form dtos.UpdateGoal) error
	AddMoneyToTotalContribute(goalId int, money float64) error
	RemoveGoalById(goalId int) error
}

type Service struct {
	Records
	Goals
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Records: NewRecordService(repository.Records),
		Goals:   NewGoalService(repository.Goals),
	}
}
