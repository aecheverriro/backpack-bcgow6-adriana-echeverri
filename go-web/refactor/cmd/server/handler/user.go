package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-web/refactor/internal/users"
	"encoding/json"
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"strings"
)

type Request struct {
	Id				string  `json:"id"`
	Name 			string	`json:"name" binding:"required"`
	LastName 		string	`json:"lastName" binding:"required"`
	Email 			string	`json:"email" binding:"required"`
	Age 			int		`json:"age" binding:"required"`
	Height 			float64 `json:"height" binding:"required"`
	Active 			bool	`json:"active" binding:"required"`
	CreationDate	bool	`json:"creationDate"`
}

type User struct {
	service users.Service
}

type Users struct {
	Users []User `json:"users"`
}

func NewUser (u users.Service) *User {
	return &User{
		service: u,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		var userQuery Request
		err := ctx.ShouldBindQuery(&userQuery)

		allUsers, err := u.service.GetAll(userQuery.Id, userQuery.Name, userQuery.lastName, userQuery.Email, userQuery.Age, userQuery.Height, userQuery.Active, userQuery.CreationDate)
		
		if err != nil {
			ctx.JSON(404, gin.H {
				"error:", err.Error()
			})
			return
		}
		ctx.JSON(200, allUsers)
		return
	}
}

func (u *User) GetId(ctx *gin.Context) {
	
	userById, err := u.service.GetId(ctx.Param("id"))

	if err != nil {
		ctx.JSON(404, gin.H {
			"error:", err.Error()
		})
		return
	}
	ctx.JSON(200, userById)
	return
}

func (u *User) CreateUser(ctx *gin.Context) {
	var request Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errorArray := strings.Split(string(err.Error()), "'")
		errorMessage := fmt.Sprintf("El campo %s es requerido", errorArray[3])
		ctx.JSON(400, gin.H{
			"error": errorMessage,
		})
		return
	}

	request.Id = fmt.Sprintf("%v", len(UsersRequest) + 1)
	request.CreationDate = "04-10-2022"
	appended := append(UsersRequest, request)
	fmt.Printf("%v, %v", appended, UsersRequest)
	ctx.JSON(200, request)
}
