package api

import (
	"hng11task2/services"
	"hng11task2/typ/jwt"
	"hng11task2/typ/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.VerifyFromBearer(c.GetHeader("Authorization"))
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userData := (claims["data"]).(map[string]interface{})
		if userData["id"] == "" {
			response.Error(c, http.StatusUnauthorized, "The authentication data is incomplete")
			c.Abort()
			return
		}

		user, err := services.GetUserById(userData["id"].(string))
		if err != nil || user == nil {
			response.Error(c, http.StatusUnauthorized, "The authenticated user may have been deleted")
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}