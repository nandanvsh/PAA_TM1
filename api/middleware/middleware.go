package middleware

import (
	"fmt"
	"foods/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsLogin(c *gin.Context) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	tokenString := cookie.Value

	token, err := service.ValidateToken(tokenString)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	userId := fmt.Sprintf("%v", token)

	c.Set("userId", userId)
	c.Next()
}
