package sorting

import (
	"errors"
)

func Sort(arrayToSort []int) ([]int, error) {
	var sortedList []int
	var minimumNumber int
	
	length := len(arrayToSort)
	if length == 0 {
		return nil, errors.New("List should not be empty")
	}

	for i := 0; i < length; i++ {
		minimumNumber = arrayToSort[0]
		position := 0
		for j := 0; j < len(arrayToSort); j++ {
			if arrayToSort[j] < minimumNumber {
				minimumNumber = arrayToSort[j]
				position = j
			}
		}
		sortedList = append(sortedList, minimumNumber)
		arrayToSort = append(arrayToSort[:position], arrayToSort[position + 1:]...)
	}
	return sortedList, nil
}
