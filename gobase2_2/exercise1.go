package main

import "fmt"

type Student struct {
	name string
	lastName string
	dni int
	startDate string
}

func (student *Student) Detail() (string) {
	dataString := fmt.Sprintf("%s, %s, %d, %s", student.name, student.lastName, student.dni, student.startDate)
	return dataString
}

func main() {
	testStudent := Student {
		"Adriana",
		"Echeverri Romero",
		10237,
		"19/08/2022",
	}

	fmt.Println(testStudent)
	fmt.Println(testStudent.Detail())
}
