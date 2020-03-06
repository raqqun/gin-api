package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/raqqun/gin-api/config"
	"github.com/raqqun/gin-api/models"
	"github.com/raqqun/gin-api/routes"
)

func main() {
	router := gin.Default()
	port := config.Port()
	// router.RedirectTrailingSlash = false

	// init database
	models.InitDB()

	// register middlewares

	// register routes
	routes.API(router)

	log.Printf("Starting on port %s", port)

	router.Run(port)
}
