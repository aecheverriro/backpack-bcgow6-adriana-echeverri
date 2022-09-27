package main

import (
	"fmt"
	"errors"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minimumCalculator(values ...float64) float64{
	minimumValue := values[0]

	for _, value :=  range values {
		if value < minimumValue {
			minimumValue = value
		}
	}

	return minimumValue
}

func maximumCalculator(values ...float64) float64{
	maximumValue := values[0]

	for _, value :=  range values {
		if value > maximumValue {
			maximumValue = value
		}
	}

	return maximumValue
}

func averageCalculator(values ...float64) float64{
	var valuesAddition float64

	for _, value :=  range values {
		valuesAddition += value
	}

	return valuesAddition / float64(len(values))
}

func operation(operator string) ((func( ...float64) float64), error) {
	switch operator {
	case minimum:
		return minimumCalculator, nil
	case average:
		return averageCalculator, nil
	case maximum:
		return maximumCalculator, nil
	}
	return nil, errors.New("The operation is not handled by the program")
}

func main() {
	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)
	
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if err == nil {
		fmt.Println(minValue)
		fmt.Println(averageValue)
		fmt.Println(maxValue)
	} else {
		fmt.Println(err)
	}
}
