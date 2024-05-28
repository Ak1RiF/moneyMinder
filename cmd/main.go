package main

import (
	"fmt"
	"time"
)

func main() {
	start := "25.05.2024"
	final := "28.05.2024"

	layot := "02.01.2006"

	t1, _ := time.Parse(layot, start)
	t2, _ := time.Parse(layot, final)

	differentDays := int((t2.Unix() - t1.Unix()) / 60 / 60 / 24)
	fmt.Println(differentDays)
}
