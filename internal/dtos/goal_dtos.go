package dtos

type CreateGoal struct {
	Title          string  `json:"title"`
	Amount         float64 `json:"amount"`
	DateCreate     string  `json:"date_create"`
	DateCompletion string  `json:"date_completion"`
}

type UpdateGoal struct {
	Title          string  `json:"title"`
	Amount         float64 `json:"amount"`
	DateCompletion string  `json:"date_completion"`
}

type GoalOutput struct {
	Id               int     `json:"id"`
	Title            string  `json:"title"`
	Amount           float64 `json:"amount"`
	DateCreate       string  `json:"date_create"`
	DateCompletion   string  `json:"date_completion"`
	DaysLeft         int     `json:"days_left"`
	RemainingAmount  float64 `json:"remaining_amount"`
	TotalContributed float64 `json:"total_contributed"`
}
