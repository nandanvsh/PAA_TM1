package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, tokenString string) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &cookie)
}

func DelCookie(c *gin.Context) {
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now(),
	}

	http.SetCookie(c.Writer, &cookie)
}
