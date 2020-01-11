package routes

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

func AuthRoutes(api *gin.RouterGroup, db *gorm.DB) {

    api.POST("/auth/signup", func (c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "ok"})
    })

    api.POST("/auth/login", func (c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "ok"})
    })
}
