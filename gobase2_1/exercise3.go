package main

import (
	"fmt"
	"errors"
)

func salaryCalculator(category string, minutes float64) (salary float64, err error) {
	var hours float64 = minutes / 60

	if minutes < 0 {
		err = errors.New("A person can't work negative time")
	}

	switch category {
	case "A":
		salary = 1000 * hours
	case "B":
		salary = 1500 * hours * 1.2
	case "C":
		salary = 3000 * hours * 1.5
	}

	return
}

func main() {
	salary, err := salaryCalculator("C", 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
}
