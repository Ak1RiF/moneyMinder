package service

import (
	"MoneyMinder/internal/dtos"
	"MoneyMinder/internal/models"
	"MoneyMinder/internal/repository"
)

type RecordService struct {
	recordRepository repository.Records
}

func NewRecordService(repo repository.Records) *RecordService {
	return &RecordService{repo}
}

func (s *RecordService) GetProfitRecords() ([]*dtos.RecordOutput, error) {
	var profits []*dtos.RecordOutput

	records, err := s.recordRepository.GetByType("profit")

	if err != nil {
		return nil, err
	}

	for _, v := range records {
		profits = append(profits, &dtos.RecordOutput{
			Id:          v.Id,
			Name:        v.Name,
			Type:        v.Type,
			Amount:      v.Amount,
			Description: v.Description,
		})

	}

	return profits, nil
}

func (s *RecordService) GetExpenseRecords() ([]*dtos.RecordOutput, error) {
	var profits []*dtos.RecordOutput

	records, err := s.recordRepository.GetByType("expenses")

	if err != nil {
		return nil, err
	}

	for _, v := range records {
		profits = append(profits, &dtos.RecordOutput{
			Id:          v.Id,
			Name:        v.Name,
			Type:        v.Type,
			Amount:      v.Amount,
			Description: v.Description,
		})

	}

	return profits, nil
}

func (s *RecordService) AddRecord(inputBody dtos.RecordInput) error {
	if err := s.recordRepository.Create(models.Record{
		Name:        inputBody.Name,
		Type:        inputBody.Type,
		Amount:      inputBody.Amount,
		Description: inputBody.Description,
	}); err != nil {
		return err
	}

	return nil
}

func (s *RecordService) UpdateRecord(recordId int, updateBody dtos.RecordInput) error {
	if err := s.recordRepository.Update(recordId, models.Record{
		Name:        updateBody.Name,
		Type:        updateBody.Type,
		Amount:      updateBody.Amount,
		Description: updateBody.Description,
	}); err != nil {
		return err
	}

	return nil
}
