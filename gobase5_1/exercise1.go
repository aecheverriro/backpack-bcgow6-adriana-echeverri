package main

import (
	"os"
)

const (
	filename = "customers.txt"
)

func main() {
	_, err := os.Open(filename)

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
}
