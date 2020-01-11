package routes

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"

    "github.com/raqqun/gin-api/models"
)

func UsersRoutes(api *gin.RouterGroup, db *gorm.DB) {

    api.GET("/users", func (c *gin.Context) {

        users := models.Users{}

        results, err := users.FindAll(db)
        if err != nil {
            c.JSON(http.StatusOK, results)
            return
        }

        c.JSON(http.StatusOK, results)
    })

    api.GET("/users/:id", func(c *gin.Context) {
        return
    })

    api.POST("/users", func (c *gin.Context) {

        user := models.Users{}
        err := c.ShouldBindJSON(&user)
        if err != nil {
            c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
            return
        }

        userCreated, err := user.Save(db)

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, userCreated)
    })

    api.PUT("/users/:id", func (c *gin.Context) {
        id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

        user := models.Users{}
        user.ID = uint(id)

        err := c.ShouldBindJSON(&user)
        if err != nil {
            c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
            return
        }

        userUpdated, err := user.Update(db)

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, userUpdated)
    })

    api.DELETE("/users/:id", func (c *gin.Context) {
        id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

        user := models.Users{}

        _, err := user.Delete(db, uint(id))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.String(http.StatusOK, "OK")
    })



}
