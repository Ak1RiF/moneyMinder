package dtos

type RecordOutput struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type RecordInput struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
