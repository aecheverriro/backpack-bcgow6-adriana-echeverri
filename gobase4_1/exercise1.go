package main

import (
	"fmt"
)

type customError struct {
	message string
}

func (err *customError) Error() string {
	return err.message
}

func main() {
	salary := 15000

	if salary < 150000 {
		err := &customError{"error: el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
