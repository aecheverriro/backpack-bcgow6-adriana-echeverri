package main

import (
	"fmt"
	"math/rand"
)

const (
	filename = "customers.txt"
)

func idGenerator() (id int, err error){
	id = rand.Intn(100000)
	return
}

func main() {
	id, err := idGenerator()

	if err != nil {
		panic("no se pudo generar un legajo adecuadamente")
	}

	fmt.Println(id)
}
