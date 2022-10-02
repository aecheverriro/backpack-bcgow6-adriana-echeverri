package main

import "fmt"

type store struct {
	products []product
}

type product struct {
	typeProduct string
	name string
	price float64
}

type Product interface {
	costCalculator() float64
}

type Ecommerce interface {
	total() float64
	add() float64
}

func newProduct(typeProduct string, name string, price float64) (product) {
	createdProduct := product {
		typeProduct,
		name,
		price,
	}
	return createdProduct
}

func (p *product) costCalculator() (float64){
	switch p.typeProduct {
	case "Small":
		return p.price * 0
	case "Medium":
		return p.price * 0.03
	case "Large":
		return p.price * 0.06
	}
	return 0
}

func (p *product) total() (float64) {
	return p.costCalculator() + p.price 
}

func (s *store) add(p product) ([]product) {
	allStores := append(s.products, p)
	return allStores
}

func newStore() (store){
	var newStore store
	return newStore
}

func main () {
	testStore := newStore()
	testProduct1 := newProduct("Large", "Table", 38.5)
	testProduct2 := newProduct("Medium", "Chair", 38.5)
	testProduct3 := newProduct("Small", "Napkins", 38.5)

	testStore.add(testProduct1)
	testStore.add(testProduct2)
	testStore.add(testProduct3)

	fmt.Printf("The price is: %v, the cost is: %v, the total is: %v\n", testProduct1.price, testProduct1.costCalculator(), testProduct1.total())
	fmt.Printf("The price is: %v, the cost is: %v, the total is: %v\n", testProduct2.price, testProduct2.costCalculator(), testProduct2.total())
	fmt.Printf("The price is: %v, the cost is: %v, the total is: %v\n", testProduct3.price, testProduct3.costCalculator(), testProduct3.total())
	
}
