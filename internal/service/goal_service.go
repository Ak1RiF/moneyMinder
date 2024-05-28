package service

import (
	"MoneyMinder/internal/dtos"
	"MoneyMinder/internal/helpers"
	"MoneyMinder/internal/models"
	"MoneyMinder/internal/repository"
)

type GoalService struct {
	goalRepository repository.Goals
}

func NewGoalService(repo repository.Goals) *GoalService {
	return &GoalService{repo}
}

func (s *GoalService) GetAllGoals() ([]*dtos.GoalOutput, error) {
	var output []*dtos.GoalOutput

	goalsFromDb, err := s.goalRepository.Get()
	if err != nil {
		return nil, err
	}

	for _, v := range goalsFromDb {
		output = append(output, &dtos.GoalOutput{
			Id:               v.Id,
			Title:            v.Title,
			Amount:           v.Amount,
			DateCreate:       v.DateCreate,
			DateCompletion:   v.DateCompletion,
			DaysLeft:         helpers.GetDifferentInDays(v.DateCreate, v.DateCompletion),
			RemainingAmount:  v.Amount - v.TotalContributed,
			TotalContributed: v.TotalContributed,
		})
	}

	return output, nil
}
func (s *GoalService) GetGoalById(id int) (*dtos.GoalOutput, error) {
	goalFromDb, err := s.goalRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return &dtos.GoalOutput{
		Id:               goalFromDb.Id,
		Title:            goalFromDb.Title,
		Amount:           goalFromDb.Amount,
		DateCreate:       goalFromDb.DateCreate,
		DateCompletion:   goalFromDb.DateCompletion,
		DaysLeft:         helpers.GetDifferentInDays(goalFromDb.DateCreate, goalFromDb.DateCompletion),
		RemainingAmount:  goalFromDb.Amount - goalFromDb.TotalContributed,
		TotalContributed: goalFromDb.TotalContributed}, nil
}
func (s *GoalService) CreateNewGoal(form dtos.CreateGoal) error {
	newGoal := models.Goal{
		Title:          form.Title,
		Amount:         form.Amount,
		DateCreate:     form.DateCreate,
		DateCompletion: form.DateCompletion,
	}
	if err := s.goalRepository.Create(newGoal); err != nil {
		return err
	}
	return nil
}
func (s *GoalService) UpdateFieldGoal(goalId int, form dtos.UpdateGoal) error {
	goalFromdb, err := s.goalRepository.GetById(goalId)
	if err != nil {
		return err
	}

	goalFromdb.Title = form.Title
	goalFromdb.Amount = form.Amount
	goalFromdb.DateCompletion = form.DateCompletion

	if err := s.goalRepository.Update(*goalFromdb); err != nil {
		return err
	}
	return nil
}
func (s *GoalService) AddMoneyToTotalContribute(goalId int, money float64) error {
	goalFromDb, err := s.goalRepository.GetById(goalId)
	if err != nil {
		return err
	}

	goalFromDb.TotalContributed += money

	if err := s.goalRepository.Update(*goalFromDb); err != nil {
		return err
	}
	return nil
}
func (s *GoalService) RemoveGoalById(goalId int) error {
	if err := s.goalRepository.Delete(goalId); err != nil {
		return err
	}
	return nil
}
