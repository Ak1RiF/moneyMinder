package models

type Goal struct {
	Id               int
	Title            string
	Amount           float64
	TotalContributed float64
	DateCreate       string
	DateCompletion   string
}
