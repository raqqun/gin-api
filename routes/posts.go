package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt "github.com/raqqun/gin-api/middleware"
)

// PostsRoutes defines posts related routes
func PostsRoutes(api *gin.RouterGroup) {

	api.POST("/posts",
		jwt.JWT().MiddlewareFunc(),
		func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{"message": "/posts", "status": http.StatusCreated})
		},
	)

	api.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "/posts", "status": http.StatusOK})
	})

	api.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		var msg struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}

		msg.Message = fmt.Sprintf("/posts/%s", id)
		msg.Status = http.StatusOK

		c.JSON(http.StatusOK, msg)
	})

}
