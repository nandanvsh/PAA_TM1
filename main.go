package main

import (
	_ "embed"
	"fmt"
	"foods/api/controller"
	"foods/api/middleware"
	"foods/api/pages"
	"foods/db"
	"foods/repo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed sql/users.sql
var users string

//go:embed sql/foods.sql
var foods string

func main() {
	database, err := db.Database("localhost", "postgres", "123", "foods", "5432")
	if err != nil {
		log.Fatalf("error connecting db : %v", err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatalf("error connecting db : %v", err)
	}

	fmt.Println("success connecting db")

	_, err = database.Exec(users)
	_, err = database.Exec(foods)

	userRepo := repo.NewUserRepo(database)
	userController := controller.NewUserController(userRepo)

	foodRepo := repo.NewFoodRepo(database)
	foodController := controller.NewFoodController(foodRepo)

	r := gin.Default()
// gin template dari golang 

	r.LoadHTMLGlob("views/*")
	r.Static("/style", "./style")

	//handler
	r.POST("/create-user", userController.CreateUser)
	r.POST("/login", userController.Login)
	r.POST("/logout", middleware.IsLogin, userController.Logout)
	r.POST("/food", middleware.IsLogin, foodController.AddFood)
	r.POST("/food/:id", middleware.IsLogin, foodController.UpdateFood)

	r.GET("/dashboard", middleware.IsLogin, foodController.GetFoodByUserId)

	//page
	r.GET("/", pages.ShowLoginPage)
	r.GET("/register", pages.ShowRegisterPage)
	r.POST("/add-food", middleware.IsLogin, pages.ShowAddFoodPage)
	r.POST("/edit-food/:id", middleware.IsLogin, func(c *gin.Context) {
		id := c.Param("id")

		food, _ := foodRepo.GetFoodById(id)
		c.HTML(http.StatusOK, "edit-food.html", food)
	})
	r.POST("/delete-food/:id", middleware.IsLogin, func(c *gin.Context) {
		foodController.DeleteFood(c)
	})

	r.Run(":8080")
}
