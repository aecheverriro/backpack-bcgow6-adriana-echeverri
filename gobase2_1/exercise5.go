package main

import "fmt"

const (
   dog = "dog"
   cat = "cat"
   hamster = "hamster"
   tarantula = "tarantula"
)

func catFoodAmount(numberCats int) int{
	return numberCats * 5000
}

func dogFoodAmount(numberDogs int) int{
	return numberDogs * 10000
}

func hamsterFoodAmount(numberHamsters int) int{
	return numberHamsters * 250
}

func tarantulaFoodAmount(numberTarantulas int) int{
	return numberTarantulas * 150
}


func Animal(animal string) ((func(int) (int)), string) {
	var returnedAnimal func(int) int = nil
	switch animal {
	case dog:
		returnedAnimal = dogFoodAmount
	case cat:
		returnedAnimal = catFoodAmount
	case hamster:
		returnedAnimal = hamsterFoodAmount
	case tarantula:
		returnedAnimal = tarantulaFoodAmount
	}

	if returnedAnimal == nil {
		return nil, "Animal does not exist"
	} else {
		return returnedAnimal, ""
	}
}

func main() {
	animalDog, msg := Animal(dog)
	animalCat, msg := Animal(cat)
	
	if msg == "" {
		var amount int
		amount += animalDog(5)
		amount += animalCat(8)
		fmt.Println(amount)
	} else {
		fmt.Println(msg)	
	}
}
 