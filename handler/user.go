package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct{
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{ "errors":errors }

		response := helper.APIResponse("Register account failed",http.StatusUnprocessableEntity,"error",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	newUser,err := h.userService.RegisterUser(input)

	if err != nil {
		// c.JSON(http.StatusBadRequest,nil)
		response := helper.APIResponse("Register account failed",http.StatusUnprocessableEntity,"error",err.Error())
		c.JSON(http.StatusBadRequest,response)

		return
	}

	// token := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser,"token")

	response := helper.APIResponse("Account has been registered",http.StatusOK,"success",formatter)

	// return newUser, nil
	c.JSON(http.StatusOK,response)
}