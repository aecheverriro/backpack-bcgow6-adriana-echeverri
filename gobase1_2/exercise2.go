package main

import "fmt"

func main() {
	var age int = 23
	var employeed bool = true
	var seniority int = 3
	var salary = 100001

	if age > 22 && employeed && seniority > 1 {
		fmt.Printf("You are eligible for a loan! \n")
		if salary > 100000 {
			fmt.Printf("With no interests!!\n")
		}
	} else {
		fmt.Printf("You are not eligible for a loan. Sorry!\n")
	}
}
