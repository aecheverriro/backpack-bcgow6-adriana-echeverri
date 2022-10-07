package users

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"errors"
	"strconv"
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
	pathToJson = "./users.json" // CHANGE TO JSON LOCATION
)

type Repository interface {
	GetAll(string, string, string, string, int, float64, bool, string) ([]User, error)
	GetId(string) (User, error)
	CreateUser(string, string, string, string, int, float64, bool, string) (User, error)
	UpdateUser(string, string, string, string, int, float64, bool) (User, error)
	DeleteUser(string) (string, error)
	PatchUser(string, string, int) (User, error)
	LastId() (string, error)
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
		// } else if active == user.Active {
		// 	filteredUsers = append(filteredUsers, user)
		} else if creation != "" && creation == user.CreationDate {
			fmt.Println("Entro!")
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
	return User{}, fmt.Errorf("Usuario %s no encontrado", id)
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

	usersJSON, err := json.MarshalIndent(allUsers, "", "  ")
	if err != nil {
		return User{}, err
	}

	err = os.WriteFile(pathToJson, usersJSON, 0644)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (r *repository) UpdateUser(id string, name string, lastName string, email string, age int, height float64, active bool) (User, error) {
	var allUsers []User
	var usersByte []byte
	jsonFile, _ := os.Open(pathToJson)
	usersByte, _ = ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(usersByte, &allUsers); err != nil {
		return User{}, err
	}

	updatedUser := User{id, name, lastName, email, age, height, active, ""}

	for index, user := range allUsers {
		if user.Id == updatedUser.Id {
			updatedUser.CreationDate = user.CreationDate
			allUsers[index] = updatedUser
			usersJSON, err := json.MarshalIndent(allUsers, "", "  ")
			if err != nil {
				return User{}, err
			}

			err = os.WriteFile(pathToJson, usersJSON, 0644)
			if err != nil {
				return User{}, err
			}
			return updatedUser, nil
		}
	}

	return User{}, errors.New("No user found with that id")
}

func (r *repository) DeleteUser(id string) (string, error) {
	var allUsers []User
	var usersByte []byte
	jsonFile, _ := os.Open(pathToJson)
	usersByte, _ = ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(usersByte, &allUsers); err != nil {
		return "", err
	}
	
	deleted := false
	var index int

	for i, user := range allUsers {
		if user.Id == id {
			deleted = true
			index = i
		}
	}

	if deleted == false {
		return "", errors.New("No user found with that id")
	}

	allUsers = append(allUsers[:index], allUsers[index+1:]...)
	usersJSON, err := json.MarshalIndent(allUsers, "", "  ")
	if err != nil {
		return "", err
	}

	err = os.WriteFile(pathToJson, usersJSON, 0644)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("El usuario %v ha sido eliminado", id), nil
}

func (r *repository) PatchUser(id string, lastName string, age int) (User, error) {
	var allUsers []User
	var usersByte []byte
	var updatedUser User
	jsonFile, _ := os.Open(pathToJson)
	usersByte, _ = ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(usersByte, &allUsers); err != nil {
		return User{}, err
	}

	for index, user := range allUsers {
		if user.Id == id {
			updatedUser.Id = id
			updatedUser.Name = user.Name
			updatedUser.LastName = lastName
			updatedUser.Email = user.Email
			updatedUser.Age = age
			updatedUser.Height = user.Height
			updatedUser.Active = user.Active
			updatedUser.CreationDate = user.CreationDate
			allUsers[index] = updatedUser

			usersJSON, err := json.MarshalIndent(allUsers, "", "  ")
			if err != nil {
				return User{}, err
			}

			err = os.WriteFile(pathToJson, usersJSON, 0644)
			if err != nil {
				return User{}, err
			}
			return updatedUser, nil
		}
	}

	return User{}, errors.New("No user found with that id")
}

func (r *repository) LastId() (string, error) {
	var allUsers []User
	var usersByte []byte
	jsonFile, _ := os.Open(pathToJson)
	usersByte, _ = ioutil.ReadAll(jsonFile)
	maxId := 0

	if err := json.Unmarshal(usersByte, &allUsers); err != nil {
		return "", err
	}

	for _, user := range allUsers {
		userId, _ := strconv.Atoi(user.Id)
		if userId > maxId {
			maxId = userId
		}
	}

	return fmt.Sprintf("%v", maxId + 1), nil
}
