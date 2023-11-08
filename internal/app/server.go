package app

import (
	"log"

	"github.com/SicParv1sMagna/HappyPetsBackend/docs"
	"github.com/SicParv1sMagna/HappyPetsBackend/internal/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (a *Application) StartServer() {
	docs.SwaggerInfo.Title = "HappyPets RestAPI"
	docs.SwaggerInfo.Description = "API server for Native HappyPets application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Создаем роутинг
	router := gin.Default()
	// Обработчик OPTIONS запросов
	router.OPTIONS("/*path", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "options",
		})
	})
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // List of allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Enable credentials (e.g., cookies)
	}))

	user := router.Group("/user")
	{
		user.POST("/register", a.handler.Register)
		user.POST("/login", a.handler.Login)
	}

	api := router.Group("/api")
	{
		api.Use(middleware.Authenticated(a.repository.GetRedisClient(), []byte("SuperSecretKey")))

		user = api.Group("/user")
		{
			user.GET("", a.handler.GetUserById)
			user.PUT("", a.handler.UpdateUserData)
		}
		pet := api.Group("pet")
		{
			image := pet.Group("/image") // Работа с изображениями минио хранилища
			{
				image.POST("/upload/:petID", a.handler.UploadImage)   // Метод для загрузки изображения
				image.DELETE("/remove/:petID", a.handler.RemoveImage) // Метод для удаления изображения
			}
			pet.POST("/create", a.handler.CreatePet)          // Метод для создания питомца
			pet.PUT("/update/:petID", a.handler.UpdatePet)    // Метод для изменения информации питомца
			pet.DELETE("/delete/:petID", a.handler.DeletePet) // Метод удаления питомца
			pet.GET("/", a.handler.GetAllPets) 			   	  // Метод вывода всех питомцев пользователя
			pet.GET("/:petID", a.handler.GetPetById)	  	  // Метод вывода питомца пользователя по petID
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run()
	if err != nil {
		log.Println("Error with running\nServer down")
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
