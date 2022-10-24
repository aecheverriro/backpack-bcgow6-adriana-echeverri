package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-testing/2_1/internal/users"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-testing/2_1/pkg/web"
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

// List Users godoc
// @Summary      Show all users
// @Description  get all users
// @Tags         Users
// @Produce      json
// @Param    token  header    int           true  "token"
// @Success  200    {object}  web.Response  "List users"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  404    {object}  web.Response  "Not found products"
// @Router   /users [GET]
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
		fmt.Println("AQUI")
		
		allUsers, err := u.service.GetAll()
		
		fmt.Println("AQUI2")

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, allUsers, ""))
	}
}

// List User godoc
// @Summary      Show user of specified id
// @Description  get user by its id
// @Tags         Users
// @Produce      json
// @Param    id       path    string        true   "ID user"
// @Param    token  header    int           true  "token"
// @Success  200    {object}  web.Response  "user"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  404    {object}  web.Response  "Not found products"
// @Router   /users/{id}      [GET]
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
}

// Store User godoc
// @Summary  Store user
// @Tags     Users
// @Accept   json
// @Produce  json
// @Param    token    header    int             true  "token requerido"
// @Param    user	  body      Request		true  "User to Store"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Router   /users   [POST]
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

	currentTime := time.Now()
	creationDate := currentTime.Format("01-02-2006")
	newUser, err := u.service.CreateUser(request.Name, request.LastName, request.Email, request.Age, request.Height, request.Active, creationDate)

	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	ctx.JSON(200, web.NewResponse(200, newUser, ""))
}

// Update User godoc
// @Summary  Update user
// @Tags     Users
// @Accept   json
// @Produce  json
// @Param    id       path      string          true   "Id user"
// @Param    token    header    int             false  "Token"
// @Param    user     body      Request         true   "User to update"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  404      {object}  web.Response
// @Router   /users/{id} [PUT]
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
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, updatedUser, ""))
}

// Update Lastname and Age godoc
// @Summary      Update last name and age of user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Description  This endpoint update field last name and age from an user
// @Param        token  header    int               true  "Token header"
// @Param        id     path      string            true  "User ID"
// @Param        name   body      Request           true  "User last name and age"
// @Success      200    {object}  web.Response
// @Failure      401    {object}  web.Response
// @Failure      400    {object}  web.Response
// @Failure      404    {object}  web.Response
// @Router       /users/{id} [PATCH]
func (u *User) PatchUser(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
		return
	}
	
	var request Request
	
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, "Hay un problema con la estructura"))
		return
	}
	
	if request.LastName == "" || request.Age == 0 {
		errorMessage := "El campo edad y apellido es requerido"
		ctx.JSON(400, web.NewResponse(400, nil, errorMessage))
		return
	}
	
	request.Id = ctx.Param("id")
	patchedUser, err := u.service.PatchUser(request.Id, request.LastName, request.Age)
	
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	ctx.JSON(200, web.NewResponse(200, patchedUser, ""))
}

// Delete User godoc
// @Summary  Delete user
// @Tags     Users
// @Param    id     path      string  true  "Product id"
// @Param    token  header    int     true  "Token"
// @Success  204
// @Failure  401    {object}  web.Response
// @Failure  404    {object}  web.Response
// @Router   /users/{id} [DELETE]
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
	ctx.JSON(204, web.NewResponse(204, message, ""))
}
