package main

import (
	"github.com/gin-gonic/gin"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-web/refactor/cmd/server/handler"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-web/refactor/internal/users"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	repo := users.NewRepository()
	service := users.NewService(repo)

	p := handler.NewUser(service)

	router := gin.Default()

	grouped := router.Group("/users")
	grouped.GET("/", p.GetAll())
	grouped.GET("/:id", p.GetId)
	grouped.POST("/", p.CreateUser)
	grouped.PUT("/:id", p.UpdateUser)
	grouped.DELETE("/:id", p.DeleteUser)
	grouped.PATCH("/:id", p.PatchUser)
	router.Run()
}
