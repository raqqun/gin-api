package routes

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    "github.com/raqqun/gin-api/models"
    "github.com/raqqun/gin-api/middleware"
)

func UsersRoutes(api *gin.RouterGroup) {

    api.GET("/users", func (c *gin.Context) {

        users := models.Users{}

        results, err := users.FindAll()
        if err != nil {
            c.JSON(http.StatusOK, results)
            return
        }

        c.JSON(http.StatusOK, results)
    })

    api.GET("/users/:id", func(c *gin.Context) {
        id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

        users := models.Users{}

        results, err := users.FindByID(uint(id))

        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, results)
    })

    api.POST("/users",
        jwt.JWT().MiddlewareFunc(),
        func (c *gin.Context) {

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

    api.PUT("/users/:id",
        jwt.JWT().MiddlewareFunc(),
        func (c *gin.Context) {
            id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

            user := models.Users{}
            user.ID = uint(id)

            err := c.ShouldBindJSON(&user)
            if err != nil {
                c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
                return
            }

            userUpdated, err := user.Update()

            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }

            c.JSON(http.StatusCreated, userUpdated)
        },
    )

    api.DELETE("/users/:id",
        jwt.JWT().MiddlewareFunc(),
        func (c *gin.Context) {
            id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

            user := models.Users{}

            _, err := user.Delete(uint(id))
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }

            c.String(http.StatusOK, "OK")
        },
    )



}
