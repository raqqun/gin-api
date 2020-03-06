package jwt

import (
	"errors"
	"log"
	"time"

	// "reflect"
	// "net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/raqqun/gin-api/models"
)

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func JWT() *jwt.GinJWTMiddleware {
	timeout, _ := time.ParseDuration("24h")

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     timeout * 10,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// Here we verify user credentials during login
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return nil, errors.New("missing Email or Password")
			}

			users := models.Users{}
			user, err := users.FindOne(&models.Users{Email: loginVals.Email}, true)

			if err != nil {
				return nil, errors.New("missing Email or Password")
			}

			if err := users.VerifyPassword(user.Password, loginVals.Password); err != nil {
				return nil, errors.New("missing Email or Password")
			}

			return user, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// Here we add additional data to the payload after authentication was validated
			if v, ok := data.(*models.Users); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			// Here we extract jwt claims and payload from token
			claims := jwt.ExtractClaims(c)

			return &models.Users{
				ID: uint(claims[identityKey].(float64)),
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// Here we check if user exists
			if v, ok := data.(*models.Users); ok {
				users := models.Users{}
				user, err := users.FindOne(&v, false)
				if err != nil {
					return false
				}

				c.Set("User", user)

				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
