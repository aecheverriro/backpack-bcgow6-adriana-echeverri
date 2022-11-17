package main

import (
	"3_2/cmd/server/handler"
	"3_2/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	cnn "3_2/db"
)

func main() {
	loadEnv()

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/products")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetByName())
	pr.GET("/:id", p.GetAll())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
