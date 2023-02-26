package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satheeshds/shortly/controllers"
	docs "github.com/satheeshds/shortly/docs"
	"github.com/satheeshds/shortly/repository"
	"github.com/satheeshds/shortly/services"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Shortner API
// @version         1.0
// @description     This is a sample url shortner server.

// @contact.name   satheesh ds
// @contact.email  davidsonsatheesh@gmail.com

// @host      localhost
// @BasePath  /api/v1

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	shortController := &controllers.ShortnerController{
		Service: &services.ShortnerService{
			Repo: repository.NewInMemoryRepository(),
		},
	}

	v1 := router.Group("/api/v1")
	{
		v1.POST("/shorten", shortController.Short)
		v1.GET("/shorten/topDomains", shortController.GetTopDomains)
	}

	router.GET("/:shortId", shortController.Redirect)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
