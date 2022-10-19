package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubstract(t *testing.T) {
	// Arrange
	num1 := 0
	num2 := 5
	expected := -5
	// Act
	result := Subtract(num1, num2)
	// Assert
	assert.Equal(t, expected, result, "El resultado fue %d, y se esperaba %d", result, expected)
}
