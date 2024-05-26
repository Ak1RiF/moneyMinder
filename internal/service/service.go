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

type Service struct {
	Records
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Records: NewRecordService(repository.Records)}
}
