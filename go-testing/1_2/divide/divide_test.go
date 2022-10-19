package divide

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestDivideZero(t *testing.T) {
	// Arrange
	denom := 0
	numerator := 15
	err := fmt.Sprintf("You can't divide by 0")
	// Act
	_, resultError := Divide(numerator, denom)
	// Assert
	assert.NotNil(t, resultError)
	assert.ErrorContains(t, resultError, err)
}

func TestDivideNormal(t *testing.T) {
	// Arrange
	denom := 5
	numerator := 15
	expected := 3
	// Act
	result, resultError := Divide(numerator, denom)
	// Assert
	assert.Equal(t, expected, result, "el resultado obtenido es: %v y el deseado es: %v", result, expected)
	assert.Nil(t, resultError)
}
