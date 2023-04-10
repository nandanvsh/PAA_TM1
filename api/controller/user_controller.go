package controller

import (
	"fmt"
	"foods/model"
	"foods/repo"
	"foods/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepo repo.UserRepo
}

func NewUserController(userRepo repo.UserRepo) *UserController {
	return &UserController{userRepo}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		errResp := fmt.Sprintf("error binding user : %v", err)
		res := model.Response{
			Message: errResp,
		}

		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.HTML(http.StatusBadRequest, "register.html", res)
		}
		return
	}

	hashedPass, err := service.HashPassword(user.Password)
	if err != nil {
		errResp := fmt.Sprintf("error hashing password : %v", err)
		res := model.Response{
			Message: errResp,
		}

		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusInternalServerError, res)
		} else {
			c.HTML(http.StatusInternalServerError, "register.html", res)
		}
		return
	}

	user.Password = hashedPass

	err = u.userRepo.CreateUser(&user)
	if err != nil {
		errResp := fmt.Sprintf("error creating user : %v", err)
		res := model.Response{
			Message: errResp,
		}

		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusInternalServerError, res)
		} else {
			c.HTML(http.StatusInternalServerError, "register.html", res)
		}
		return
	}

	res := model.DataResponse{
		Data:    user,
		Message: "success creating user",
	}

	if c.Request.Header.Get("Accept") == "application/json" {
		c.JSON(http.StatusCreated, res)
	} else {
		c.HTML(http.StatusCreated, "login.html", res)
	}
}

func (u *UserController) Login(c *gin.Context) {
	var user model.LoginRequest
	err := c.ShouldBind(&user)
	if err != nil {
		res := model.Response{
			Message: "error binding user",
		}
		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.HTML(http.StatusBadRequest, "login.html", res)
		}
		return
	}

	userResp, err := u.userRepo.GetUserByUsername(user.Username)
	if err != nil {
		res := model.Response{
			Message: "user not found",
		}
		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.HTML(http.StatusBadRequest, "login.html", res)
		}
		return
	}

	err = service.ValidatePassword(user.Password, userResp.Password)
	if err != nil {
		res := model.Response{
			Message: "wrong password",
		}
		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.HTML(http.StatusBadRequest, "login.html", res)
		}
		return
	}

	tokenStr, err := service.GenerateToken(userResp)
	if err != nil {
		res := model.Response{
			Message: "error generate token",
		}
		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.HTML(http.StatusBadRequest, "login.html", res)
		}
		return
	}

	service.SetCookie(c, tokenStr)

	if c.Request.Header.Get("Accept") == "application/json" {
		res := model.DataResponse{
			Data:    tokenStr,
			Message: "success login",
		}
		c.JSON(http.StatusOK, res)
	} else {
		c.Redirect(http.StatusFound, "/dashboard")
	}
}

func (u *UserController) Logout(c *gin.Context) {
	service.DelCookie(c)

	res := model.DataResponse{
		Data:    nil,
		Message: "success logout",
	}

	if c.Request.Header.Get("Accept") == "application/json" {
		c.JSON(http.StatusOK, res)
	} else {
		c.HTML(http.StatusOK, "login.html", res)
	}
}
