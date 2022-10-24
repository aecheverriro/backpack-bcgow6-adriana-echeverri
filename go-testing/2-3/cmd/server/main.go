package main

import (
	"github.com/gin-gonic/gin"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-testing/2_1/cmd/server/handler"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-testing/2_1/internal/users"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-testing/2_1/pkg/store"
	"github.com/aecheverriro/backpack-bcgow6-adriana-echeverri/go-testing/2_1/docs"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title           MELI Bootcamp User - API
// @version         1.0
// @description     This is a simple API that manages users in an ecommerce page.
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name   Adriana Echeverri
// @contact.url    http://www.swagger.io/support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	p := handler.NewUser(service)

	router := gin.Default()
	api := router.Group("/api/v1")
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	grouped := router.Group("/users")
	grouped.GET("/", p.GetAll())
	grouped.GET("/:id", p.GetId)
	grouped.POST("/", p.CreateUser)
	grouped.PUT("/:id", p.UpdateUser)
	grouped.DELETE("/:id", p.DeleteUser)
	grouped.PATCH("/:id", p.PatchUser)
	err = router.Run()
	if err != nil {
		log.Fatal("error al levantar la API")
	}
}
