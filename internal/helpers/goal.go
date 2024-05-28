package helpers

import "time"

// work with dates of goals

func GetDifferentInDays(start, final string) int {
	date_layout := "02.01.2006"

	t1, _ := time.Parse(date_layout, start)
	t2, _ := time.Parse(date_layout, final)

	return int((((t2.Unix() - t1.Unix()/60) / 60) / 24))
}

// work with amount money of goals

func GetRemainingAmount(amount, totalContrib float64) float64 {
	return amount - totalContrib
}
