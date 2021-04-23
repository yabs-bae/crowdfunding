package main

import (
	"crowdfunding/handler"
	"crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	api := router.Group("/api/v1")
	api.POST("users", userHandler.RegisterUser)


	router.Run()



	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Alfian"
	// userInput.Email = "alfianprast@gmail.com"
	// userInput.Occupation = "Project Manager"
	// userInput.PasswordHash = "password"

	// userService.RegisterUser(userInput)


}
