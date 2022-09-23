package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var amountOlder21 int = 0

	for _, age := range employees {
		if age > 21 {
			amountOlder21 += 1
		}
	}

	fmt.Printf("Benjamin age is: %v\n", employees["Benjamin"])
	fmt.Printf("The amount of employees > 21 is: %v\n", amountOlder21)
	employees["Federico"] = 25
	delete(employees, "Pedro")
}
