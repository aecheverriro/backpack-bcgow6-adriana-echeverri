package users

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type User struct {
	Id 				string	`json:"id"`
	Name 			string	`json:"name"`
	LastName 		string	`json:"lastName"`
	Email 			string	`json:"email"`
	Age 			int		`json:"age"`
	Height 			float64 `json:"height"`
	Active 			bool	`json:"active"`
	CreationDate 	string	`json:"creationDate"`
}

const (
	pathToJson = "users.json" // CHANGE TO JSON LOCATION
)

type Repository interface {
	GetAll(string, string, string, string, int, float64, bool, string) ([]User, error)
	GetId(string) (User, error)
	CreateUser(string, string, string, string, int, float64, bool, string) (User, error)
}

type repository struct {}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll(id string, name string, lastName string, email string, age int, height float64, active bool, creation string) ([]User, error) {

	var allUsers []User
	var filteredUsers []User
	var usersByte []byte
	jsonFile, _ := os.Open(pathToJson)
	usersByte, _ = ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(usersByte, &allUsers); err != nil {
		return nil, err
	}
	
	for _, user := range allUsers {
		if id != "" && id == user.Id {
			filteredUsers = append(filteredUsers, user)
		} else if name != "" && name == user.Name {
			filteredUsers = append(filteredUsers, user)
		} else if lastName != "" && lastName == user.LastName {
			filteredUsers = append(filteredUsers, user)
		} else if email != "" && email == user.Email {
			filteredUsers = append(filteredUsers, user)
		} else if age != 0 && age == user.Age {
			filteredUsers = append(filteredUsers, user)
		} else if height != 0.0 && height == user.Height {
			filteredUsers = append(filteredUsers, user)
		} else if active == user.Active {
			filteredUsers = append(filteredUsers, user)
		} else if creation != "" && creation == user.CreationDate {
			filteredUsers = append(filteredUsers, user)
		}
	}

	if len(filteredUsers) == 0 {
		return allUsers, nil
	} else {
		return filteredUsers, nil
	}
}

func (r *repository) GetId(id string) (User, error) {
	var allUsers []User
	var usersByte []byte
	jsonFile, _ := os.Open(pathToJson)
	usersByte, _ = ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(usersByte, &allUsers); err != nil {
		return User{}, err
	}

	for _, user := range allUsers {
		if user.Id == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("Usuario %d no encontrado", id)
}

func (r *repository) CreateUser(id string, name string, lastName string, email string, age int, height float64, active bool, creationDate string) (User, error) {
	var allUsers []User
	var usersByte []byte
	jsonFile, _ := os.Open(pathToJson)
	usersByte, _ = ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(usersByte, &allUsers); err != nil {
		return User{}, err
	}

	newUser := User{id, name, lastName, email, age, height, active, creationDate}

	allUsers = append(allUsers, newUser)

	usersJSON, _ := json.Marshal(allUsers)
	_, err := jsonFile.Write(usersJSON)

	if err != nil {
		return User{}, err
	}
	return newUser, nil
}
