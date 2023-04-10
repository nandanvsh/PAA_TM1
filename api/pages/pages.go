package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func ShowAddFoodPage(c *gin.Context) {
	c.HTML(http.StatusOK, "add-food.html", nil)
}
