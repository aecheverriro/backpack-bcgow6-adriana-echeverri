package main

import (
	"fmt"
	"os"
)

const (
	fileName = "./register.csv"
)

func registerSale(id string, price float64, quantity int) (err error) {
	priceFormatted := fmt.Sprintf("%f", price)
	quantityFormatted := fmt.Sprintf("%d", quantity)
	content := id + ";" + priceFormatted + ";" + quantityFormatted
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	os.Chmod(fileName, 0777)
	_, err = f.WriteString(content)
	return
}

func main() {
	id := "G15"
	price := 1.2
	quantity := 15

	err := registerSale(id, price, quantity)
	if err != nil {
		fmt.Println(err)
	}
}
if err != nil {
    panic(err)
}
