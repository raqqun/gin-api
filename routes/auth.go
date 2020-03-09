package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	jwt "github.com/raqqun/gin-api/middleware"
	"github.com/raqqun/gin-api/models"
)

// AuthRoutes defines authentication related routes
func AuthRoutes(api *gin.RouterGroup) {

	api.POST("/auth/login", jwt.JWT().LoginHandler)
	api.GET("/auth/refresh", jwt.JWT().RefreshHandler)

	api.POST("/auth/signup",
		func(c *gin.Context) {
			user := models.Users{}
			err := c.ShouldBindJSON(&user)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return
			}

			userCreated, err := user.Save()

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, userCreated)
		},
	)
}
