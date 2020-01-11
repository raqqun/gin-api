package routes

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

// here we initialize all routes

func API(r *gin.Engine, db *gorm.DB) {
    api := r.Group("/api")
    {
        api.GET("/", func (c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "/api", "status": http.StatusOK})
        })

        AuthRoutes(api, db)
        UsersRoutes(api, db)
        PostsRoutes(api, db)
    }
}
