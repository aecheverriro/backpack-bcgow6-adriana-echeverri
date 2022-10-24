package users

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

type StubStoreEngine struct{
	mockedData []User
	mockedReadCalled bool
}

func (s *StubStoreEngine) ReadUser(data interface{}) error {
	s.mockedReadCalled = true
	a, ok := data.(*[]User)
	if !ok {
		return errors.New("Failed converting")
	}
	*a = s.mockedData
	return nil
}

func (s *StubStoreEngine) WriteUser(data interface{}) error {
	a, ok := data.([]User)
	if !ok {
		return errors.New("Failed converting")
	}
	s.mockedData = append(s.mockedData, a[len(a) - 1])
	return nil
}

func TestGetAll(t *testing.T) {
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

	//Act
	result, _ := motor.GetAll()

	//Assert
	assert.Equal(t, result, expected, "The result obtained is: %v, and the expected is: %v", result, expected)
}

func TestPatchName(t *testing.T) {
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
		Name: "Adriana",
		LastName: "After Update",
		Email: "adriana.er98@hotmail.com",
		Age: 50,
		Height: 1.7,
		Active: true,
		CreationDate: "03-10-2022",
	}
	
	myStubSearchEngine := StubStoreEngine{
		mockedData: data,
		mockedReadCalled: false,
	}
	motor := NewRepository(&myStubSearchEngine)

	//Act
	result, _ := motor.PatchUser("1234", "After Update", 50)

	//Assert
	assert.Equal(t, result, dataUpdated, "The result obtained is: %v, and the expected is: %v", result, dataUpdated)
	assert.True(t, myStubSearchEngine.mockedReadCalled)
}
