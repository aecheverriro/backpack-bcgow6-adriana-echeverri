package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"strings"
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

type Users struct {
	Users []User `json:"users"`
}

type Request struct {
	Id 				string	`json:"id"`
	Name 			string	`json:"name" binding:"required"`
	LastName 		string	`json:"lastName" binding:"required"`
	Email 			string	`json:"email" binding:"required"`
	Age 			int		`json:"age" binding:"required"`
	Height 			float64 `json:"height" binding:"required"`
	Active 			bool	`json:"active" binding:"required"`
	CreationDate 	string	`json:"creationDate"`
}

const (
	pathToJson = "users.json"
)

var UsersRequest []Request

func GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var AllUsers Users
		var usersByte []byte
		jsonFile, _ := os.Open(pathToJson)
		usersByte, _ = ioutil.ReadAll(jsonFile)

		if err := json.Unmarshal(usersByte, &AllUsers); err != nil {
			log.Fatal(err)
		}
		
		var userQuery User
		err := ctx.ShouldBindQuery(&userQuery)
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println(userQuery)
		var filteredUsers []User
		for _, user := range AllUsers.Users {
			if userQuery.Id != "" && userQuery.Id == user.Id {
				filteredUsers = append(filteredUsers, user)
			} else if userQuery.Name != "" && userQuery.Name == user.Name {
				filteredUsers = append(filteredUsers, user)
			} else if userQuery.LastName != "" && userQuery.LastName == user.LastName {
				filteredUsers = append(filteredUsers, user)
			} else if userQuery.Email != "" && userQuery.Email == user.Email {
				filteredUsers = append(filteredUsers, user)
			} else if userQuery.Age != 0 && userQuery.Age == user.Age {
				filteredUsers = append(filteredUsers, user)
			} else if userQuery.Height != 0.0 && userQuery.Height == user.Height {
				filteredUsers = append(filteredUsers, user)
			} else if userQuery.Active == user.Active {
				filteredUsers = append(filteredUsers, user)
			} else if userQuery.CreationDate != "" && userQuery.CreationDate == user.CreationDate {
				filteredUsers = append(filteredUsers, user)
			}
		}

		if len(filteredUsers) == 0 {
			ctx.JSON(200, AllUsers.Users)
		} else {
			ctx.JSON(200, filteredUsers)
		}

	}
}

func GetId(ctx *gin.Context) {
	var AllUsers Users
	var usersByte []byte
	jsonFile, _ := os.Open(pathToJson)
	usersByte, _ = ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(usersByte, &AllUsers); err != nil {
		log.Fatal(err)
	}

	for _, user := range AllUsers.Users {
		if user.Id == ctx.Param("id") {
			ctx.JSON(200, user)
			return
		}
	}

	ctx.JSON(404, nil)
}

func CreateUser(ctx *gin.Context) {
	var request Request

	token := ctx.GetHeader("token")
	if token != "1234" {
		ctx.JSON(400, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errorArray := strings.Split(string(err.Error()), "'")
		errorMessage := fmt.Sprintf("El campo %s es requerido", errorArray[3])
		ctx.JSON(400, gin.H{
			"error": errorMessage,
		})
		return
	}

	request.Id = fmt.Sprintf("%v", len(UsersRequest) + 1) // ARREGLAR PARA QUE SEA ULTIMO + 1
	request.CreationDate = "04-10-2022"
	UsersRequest = append(UsersRequest, request)
	fmt.Printf("%v", UsersRequest)
	ctx.JSON(200, request)
}

func UpdateUser(ctx *gin.Context) {
	var request Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errorArray := strings.Split(string(err.Error()), "'")
		errorMessage := fmt.Sprintf("El campo %s es requerido", errorArray[3])
		ctx.JSON(400, gin.H{
			"error": errorMessage,
		})
		return
	}

	for index, user := range UsersRequest {
		if user.Id == ctx.Param("id") {
			request.Id = user.Id
			request.CreationDate = user.CreationDate
			UsersRequest[index] = request

			ctx.JSON(200, request)
			return
		}
	}

	ctx.JSON(404, nil)
}

func DeleteUser(ctx *gin.Context) {
	deleted := false
	var index int

	for i, user := range UsersRequest {
		if user.Id == ctx.Param("id") {
			deleted = true
			index = i
		}
	}

	if deleted == false {
		ctx.JSON(404, nil)
		return
	}

	UsersRequest = append(UsersRequest[:index], UsersRequest[index+1:]...)
	ctx.JSON(200, gin.H{ "data": fmt.Sprintf("El usuario %v ha sido eliminado", ctx.Param("id"))})
	return
}

func PatchUser(ctx *gin.Context) {

	var request Request

	ctx.ShouldBindJSON(&request)

	if request.LastName == "" || request.Age == 0 {
		errorMessage := "El campo edad y apellido es requerido"
		ctx.JSON(400, gin.H{
			"error": errorMessage,
		})
		return
	}

	for index, user := range UsersRequest {
		if user.Id == ctx.Param("id") {
			request.Id = user.Id
			request.Name = user.Name
			request.Email = user.Email
			request.Height = user.Height
			request.Active = user.Active
			request.CreationDate = user.CreationDate
			UsersRequest[index] = request

			ctx.JSON(200, request)
			return
		}
	}

	ctx.JSON(404, nil)
}

func main() {
	router := gin.Default()
	
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Adriana",
		})
	})

	router.GET("/users", GetAll())
	router.GET("/users/:id", GetId)
	router.POST("/users", CreateUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)
	router.PATCH("/users/:id", PatchUser)
	

	router.Run()
}

// Se requiere implementar una funcionalidad que modifique la entidad parcialmente, solo se
// deben modificar 2 campos:
// - Si se seleccionó Productos, los campos nombre y precio.
// - Si se seleccionó Usuarios, los campos apellido y edad.
// - Si se seleccionó Transacciones, los campos código de transacción y monto.
// .Para lograrlo, es necesario, seguir los siguientes pasos:
// 1. Generar un método PATCH para modificar la entidad parcialmente, modificando solo 2
// campo (a elección)
// 2. Desde el Path enviar el ID de la entidad que se modificara
// 3. En caso de no existir, retornar un error 404
// 4. Realizar validaciones de ambos campos