package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestSortingEmpty(t *testing.T) {
	// Arrange
	var emptyArray []int
	err := fmt.Sprintf("List should not be empty")
	// Act
	_, resultError := Sort(emptyArray)
	// Assert
	assert.NotNil(t, resultError)
	assert.ErrorContains(t, resultError, err)
}

func TestSortingNormal(t *testing.T) {
	// Arrange
	testArray := []int{9, 450, 56, 3, 5, 60, 199, -10}
	expectedArray := []int{-10, 3, 5, 9, 56, 60, 199, 450}
	// Act
	result, resultError := Sort(testArray)
	// Assert
	assert.Equal(t, expectedArray, result, "el arreglo obtenido es: %v y el deseado es: %v", result, expectedArray)
	assert.Nil(t, resultError)
}
