package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	fileName = "./register.csv"
)

func printRegisters() (err error) {
	data, err := os.ReadFile(fileName)
	rows := strings.Split(string(data[:]), "\n")

	fmt.Printf("ID\tPrecio\t\tCantidad\n")
	
	for _, row := range rows {
		columns := strings.Split(string(row[:]), ";")
		for _, column := range columns {
			fmt.Printf("%s\t", fmt.Sprint(column))
		}
		fmt.Printf("\n")
	}

	return
}

func main() {

	err := printRegisters()
	if err != nil {
		fmt.Println(err)
	}
}
