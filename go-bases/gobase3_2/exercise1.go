package main

import "fmt"

type User struct {
	name string
	lastName string
	age int
	email string
	password string
}

func (u *User) changeName(name string) (string){
	u.name = name
	return "ok"
}

func (u *User) changeAge(age int) (string){
	u.age = age
	return "ok"
}

func (u *User) changeEmail(email string) (string){
	u.email = email
	return "ok"
}

func (u *User) changePassword(password string) (string){
	u.password = password
	return "ok"
}

func main() {
	testUser := User {
		"Ardiana",
		"Echeverri Romero",
		23,
		"adrianaer98@hotmail.com",
		"passwordTest",
	}

	fmt.Println(testUser)
	
	testUser.changeAge(24)
	testUser.changeEmail("adriana.er98@hotmail.com")
	testUser.changeName("Adriana")
	testUser.changePassword("qwerty")
	
	fmt.Println(testUser)
}
