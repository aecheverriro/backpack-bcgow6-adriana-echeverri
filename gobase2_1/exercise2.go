package main

import (
	"fmt"
	"errors"
)

func averageCalculator(grades ...float64) (average float64, err error) {
	amountGrades := len(grades)

	if amountGrades == 0 {
		return
	}
	for _, value := range grades {
		if (value < 0) {
			err = errors.New("Grades should not be negative")
			average = 0
			return
		}
		average += value
	}

	average /= float64(len(grades))
	return
}

func main() {
	average1, _ := averageCalculator(97, 98, 98, 99, 100, 100)
	average2, _ := averageCalculator(97)
	average3, _ := averageCalculator()
	average4, err4 := averageCalculator(90, -97)
	fmt.Println(average1, average2, average3, average4)
	fmt.Println(err4)
}