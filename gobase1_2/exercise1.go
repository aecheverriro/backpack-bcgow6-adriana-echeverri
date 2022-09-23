package main

import "fmt"

func main() {
	word := "word test"
	var length int = len(word)
	index := 0

	fmt.Printf("The length of the word is:%v\n", length)

	for index < length {
		fmt.Printf("%c\n", word[index])
		index++
	}
}
