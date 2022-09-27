package main

import (
	"fmt"
	"errors"
)

func taxCalculator (salary float64) (tax float64, err error) {
	if salary < 0 {
		err = errors.New("Salary must be greater than 0")
		return
	}
	if salary > 50000 {
		tax += salary * 0.17
	}
	if salary > 150000 {
		tax += salary * 0.1
	}
	return
}

func main() {
	tax, err := taxCalculator(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tax)
	}
}
