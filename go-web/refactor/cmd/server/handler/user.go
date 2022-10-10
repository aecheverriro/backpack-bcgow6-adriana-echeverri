package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-web/refactor/internal/users"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-web/refactor/pkg/web"
	"fmt"
	"strings"
	"time"
	"os"
)

type Request struct {
	Id 				string	`json:"id"`
	Name 			string	`json:"name"`
	LastName 		string	`json:"lastName" binding:"required"`
	Email 			string	`json:"email" binding:"required"`
	Age 			int		`json:"age" binding:"required"`
	Height 			float64 `json:"height" binding:"required"`
	Active 			bool	`json:"active" binding:"required"`
	CreationDate 	string	`json:"creationDate"`
}

type User struct {
	service users.Service
}

type Users struct {
	Users []User
}

func NewUser (u users.Service) *User {
	return &User{
		service: u,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}
		
		var userQuery Request
		err := ctx.ShouldBindQuery(&userQuery)

		fmt.Println(userQuery)

		allUsers, err := u.service.GetAll(userQuery.Id, userQuery.Name, userQuery.LastName, userQuery.Email, userQuery.Age, userQuery.Height, userQuery.Active, userQuery.CreationDate)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, allUsers, ""))
		return
	}
}

func (u *User) GetId(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
		return
	}

	userById, err := u.service.GetId(ctx.Param("id"))

	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, userById, ""))
	return
}

func (u *User) CreateUser(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
		return
	}
	
	var request Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errorArray := strings.Split(string(err.Error()), "'")
		errorMessage := fmt.Sprintf("El campo %s es requerido", errorArray[3])
		ctx.JSON(400, web.NewResponse(400, nil, errorMessage))
		return
	}

	id := "4"
	currentTime := time.Now()
	creationDate := currentTime.Format("01-02-2006")
	newUser, err := u.service.CreateUser(id, request.Name, request.LastName, request.Email, request.Age, request.Height, request.Active, creationDate)

	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	ctx.JSON(200, web.NewResponse(200, newUser, ""))
	return
}

func (u *User) UpdateUser(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
		return
	}
	
	var request Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	request.Id = ctx.Param("id")
	updatedUser, err := u.service.UpdateUser(request.Id, request.Name, request.LastName, request.Email, request.Age, request.Height, request.Active)

	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, updatedUser, ""))
	return
}

func (u *User) PatchUser(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
		return
	}
	
	var request Request
	
	ctx.ShouldBindJSON(&request)
	
	if request.LastName == "" || request.Age == 0 {
		errorMessage := "El campo edad y apellido es requerido"
		ctx.JSON(400, web.NewResponse(400, nil, errorMessage))
		return
	}
	
	request.Id = ctx.Param("id")
	patchedUser, err := u.service.PatchUser(request.Id, request.LastName, request.Age)
	
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, patchedUser, ""))
	return
}

func (u *User) DeleteUser(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
		return
	}
	
	message, err := u.service.DeleteUser(ctx.Param("id"))
	
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, message, ""))
	return
}
