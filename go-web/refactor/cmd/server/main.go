package main

import (
	"github.com/gin-gonic/gin"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-web/refactor/cmd/server/handler"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-web/refactor/internal/users"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)

	p := handler.NewUser(service)

	r := gin.Default()

	
	router.GET("/users", GetAll())
	router.GET("/users/:id", GetId)
	router.POST("/users", CreateUser)
	r.Run()
}