package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nurul-iman-blok-m/auth"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/user"
)

type userHandler struct {
	userService user.UserService
	authService auth.Service
}

func NewUserHandler(userService user.UserService, authService auth.Service) *userHandler {
	return &userHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// next for add user must be super admin

	userInput, roleName, errInput := h.userService.RegisterUser(input)
	if errInput != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, errToken := h.authService.GenerateToken(userInput.ID)
	if errToken != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.UserJsonFormatter(userInput, roleName, token)

	response := helper.ApiResponse("Account has been register", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, roleName, errLogin := h.userService.LoginUser(input)
	if errLogin != nil {
		errorMessage := gin.H{"errors": errLogin.Error()}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, errToken := h.authService.GenerateToken(loggedInUser.ID)
	if errToken != nil {
		response := helper.ApiResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.UserJsonFormatter(loggedInUser, roleName, token)
	response := helper.ApiResponse("Login Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
