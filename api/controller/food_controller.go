package controller

import (
	"fmt"
	"foods/model"
	"foods/repo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FoodController struct {
	foodRepo repo.FoodRepo
}

func NewFoodController(foodRepo repo.FoodRepo) *FoodController {
	return &FoodController{foodRepo}
}

func (f *FoodController) GetFoodByUserId(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	foods, err := f.foodRepo.GetFoodByUserId(userId)
	if err != nil {
		errResp := fmt.Sprintf("error getting all foods : %v", err)
		res := model.Response{
			Message: errResp,
		}

		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusInternalServerError, res)
		} else {
			c.HTML(http.StatusInternalServerError, "error.tmpl", res)
		}
		return
	}

	res := model.DataResponse{
		Data:    foods,
		Message: "success getting all book",
	}

	if c.Request.Header.Get("Accept") == "application/json" {
		c.JSON(http.StatusOK, res)
	} else {
		c.HTML(http.StatusOK, "dashboard.html", res)
	}
	return
}

func (f *FoodController) AddFood(c *gin.Context) {
	var food model.Food
	userId := c.MustGet("userId").(string)
	userIdInt, err := strconv.Atoi(userId)

	err = c.ShouldBind(&food)
	if err != nil {
		res := model.Response{
			Message: "error binding user",
		}
		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.HTML(http.StatusBadRequest, "add-food.html", res)
		}
		return
	}

	food.User_ID = userIdInt
	err = f.foodRepo.AddFood(&food)
	if err != nil {
		errResp := fmt.Sprintf("error adding new food : %v", err)
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
		Data:    food,
		Message: "success adding new food",
	}

	if c.Request.Header.Get("Accept") == "application/json" {
		c.JSON(http.StatusCreated, res)
	} else {
		f.GetFoodByUserId(c)
	}
}

func (f *FoodController) UpdateFood(c *gin.Context) {
	param := c.Param("id")
	var food model.Food
	err := c.ShouldBind(&food)

	paramint, _ := strconv.Atoi(param)
	if err != nil {
		res := model.Response{
			Message: "error binding user",
		}
		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.HTML(http.StatusBadRequest, "edit-food.html", res)
		}
		return
	}

	err = f.foodRepo.UpdateFood(param, food)
	if err != nil {
		res := model.Response{
			Message: "error updating food",
		}
		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusBadRequest, res)
		} else {
			c.HTML(http.StatusBadRequest, "edit-food.html", res)
		}
		return
	}

	food.ID = paramint

	res := model.DataResponse{
		Data:    food,
		Message: "success updating food",
	}

	if c.Request.Header.Get("Accept") == "application/json" {
		c.JSON(http.StatusCreated, res)
	} else {
		f.GetFoodByUserId(c)
	}
}

func (f *FoodController) DeleteFood(c *gin.Context) {
	param := c.Param("id")

	food, err := f.foodRepo.GetFoodById(param)
	if err != nil {
		errResp := fmt.Sprintf("error getting food : %v", err)
		res := model.Response{
			Message: errResp,
		}

		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusInternalServerError, res)
		} else {
			c.HTML(http.StatusInternalServerError, "error.tmpl", res)
		}
		return
	}

	err = f.foodRepo.DeleteFood(param)
	if err != nil {
		errResp := fmt.Sprintf("error deleting food : %v", err)
		res := model.Response{
			Message: errResp,
		}

		if c.Request.Header.Get("Accept") == "application/json" {
			c.JSON(http.StatusInternalServerError, res)
		} else {
			c.HTML(http.StatusInternalServerError, "error.tmpl", res)
		}
		return
	}

	res := model.DataResponse{
		Data:    food,
		Message: "success deleting food",
	}

	if c.Request.Header.Get("Accept") == "application/json" {
		c.JSON(http.StatusOK, res)
	} else {
		f.GetFoodByUserId(c)
	}
}



