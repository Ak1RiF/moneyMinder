package models

// type : Profit or Expenses

type Record struct {
	Id          int
	Name        string
	Type        string
	Amount      float64
	Description string
}
