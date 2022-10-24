package products

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBySeller(t *testing.T) {
	//Arrange
	expected := []Product{{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	}}
	
	motor := NewRepository()
	service := NewService(motor)

	//Act
	result, err := service.GetAllBySeller("123")

	//Assert
	assert.Equal(t, result, expected, "The result obtained is: %v, and the expected is: %v", result, expected)
	assert.Nil(t, err)
}
