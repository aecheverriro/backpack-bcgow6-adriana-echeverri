package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	values [][]float64
	height int
	width int
	cuadratic bool
	maximum float64
}


func (m *Matrix) Set(values ...float64) (error) {
	if len(values) != m.height * m.width {
		return errors.New("Not enough values declared")
	}

	filledColumn := 0
	filledRows := 0
	for index, value := range values {
		m.values[filledRows][filledColumn] = value
		filledColumn++

		if (index + 1) % (m.width) == 0 {
			filledRows++
			filledColumn = 0
		}
	}
	return nil
}

func (m *Matrix) Print() {

	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			fmt.Printf("%v ", m.values[i][j])
		}
		fmt.Println()
	}
}

func main () {
	valuesTest := [][]float64{
		{0.1, 0.1},
		{0.1, 0.1},
	}
	testMatrix := Matrix {
		valuesTest,
		2,
		2,
		false,
		0,
	}

	testMatrix.Set(1.5, 1.2, 1.3, 1.7)
	testMatrix.Print()
}
