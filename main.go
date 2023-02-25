package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satheeshds/shortly/controllers"
	"github.com/satheeshds/shortly/repository"
	"github.com/satheeshds/shortly/services"
)

func main() {
	router := gin.Default()

	shortController := &controllers.ShortnerController{
		Service: &services.ShortnerService{
			Repo: repository.NewInMemoryRepository(),
		},
	}

	router.GET("/shorten/:shortId", shortController.Redirect)
	router.POST("/shorten", shortController.Short)
	router.GET("/shorten/topDomains", shortController.GetTopDomains)

	router.Run(":8080")
}
