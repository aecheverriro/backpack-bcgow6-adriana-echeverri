package users

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	//Arrange
	expected := []User{
		{
		  Id: "1234",
		  Name: "Adriana",
		  LastName: "Echeverri",
		  Email: "adriana.er98@hotmail.com",
		  Age: 23,
		  Height: 1.7,
		  Active: true,
		  CreationDate: "03-10-2022",
		},
		{
			Id: "12345",
			Name: "Jhon",
			LastName: "Ramirez",
			Email: "jhon.r@meli.com",
			Age: 51,
			Height: 1.9,
			Active: true,
			CreationDate: "03-10-2022",
		},
	}
	myStubSearchEngine := StubStoreEngine{
		mockedData: expected,
		mockedReadCalled: false,
	}
	motor := NewRepository(&myStubSearchEngine)
	service := NewService(motor)

	//Act
	result, err := service.DeleteUser("12345")
	expectedMessage := "El usuario 12345 ha sido eliminado"

	//Assert
	assert.Equal(t, result, expectedMessage, "The result obtained is: %v, and the expected is: %v", result, expectedMessage)
	assert.True(t, myStubSearchEngine.mockedReadCalled)
	assert.Nil(t, err)
}

func TestUpdateName(t *testing.T) {
	//Arrange
	data := []User{
		{
			Id: "1234",
			Name: "Adriana",
			LastName: "Before Update",
			Email: "adriana.er98@hotmail.com",
			Age: 23,
			Height: 1.7,
			Active: true,
			CreationDate: "03-10-2022",
		},
	}
	dataUpdated := User{
		Id: "1234",
		Name: "Mary",
		LastName: "After Update",
		Email: "adriana@gmail.com",
		Age: 25,
		Height: 1.69,
		Active: false,
		CreationDate: "03-10-2022",
	}
	
	myStubSearchEngine := StubStoreEngine{
		mockedData: data,
		mockedReadCalled: false,
	}
	motor := NewRepository(&myStubSearchEngine)
	service := NewService(motor)

	//Act
	result, err := service.UpdateUser("1234", "Mary", "After Update", "adriana@gmail.com", 25, 1.69, false)

	//Assert
	assert.Equal(t, result, dataUpdated, "The result obtained is: %v, and the expected is: %v", result, dataUpdated)
	assert.Nil(t, err)
	assert.True(t, myStubSearchEngine.mockedReadCalled)
}
