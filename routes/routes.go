package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/raqqun/gin-api/middleware"
)

// here we initialize all routes

func API(r *gin.Engine) {
	api := r.Group("/api")
	{

		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "/api", "status": http.StatusOK})
		})

		AuthRoutes(api)

		// api.Use(jwt.JWT().MiddlewareFunc())

		UsersRoutes(api)
		PostsRoutes(api)
	}
}
