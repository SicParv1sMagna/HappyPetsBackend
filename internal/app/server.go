package app

import (
	"log"

	"github.com/SicParv1sMagna/HappyPetsBackend/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (a *Application) StartServer() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Создаем роутинг
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // List of allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Enable credentials (e.g., cookies)
	}))

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", a.handler.Register)
		}
		image := api.Group("/image")//Работа с изображениями минио хранилища
		{
			image.POST("/upload/:userID/:petID", a.handler.UploadImage) // Метод для загрузки изображения
			image.DELETE("/remove/:userID/:petID", a.handler.RemoveImage) // Метод для удаления изображения
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run()
	if err != nil {
		log.Println("Error with running\nServer down")
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
