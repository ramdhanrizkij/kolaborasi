package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kolaborasi/dto"
	"kolaborasi/entity"
	"kolaborasi/helper"
	"kolaborasi/service"
	"net/http"
)

type userHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserHandler(userService service.UserService, authService service.AuthService) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input dto.RegisterUserDto

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("error", "Please Check Your Data", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Cek apakah email sudah terdaftar
	emailInput := dto.EmailCheckDTO{
		Email: input.Email,
	}
	available, err := h.userService.IsEmailAvailable(emailInput)
	fmt.Printf("Format email %v", available)
	if !available {
		response := helper.APIResponse("error", "Registration User Failed, Email already registered", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//Save user data
	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("error", "Registration User Failed", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error", "Login Failed, Failed to generate JWT Token", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userFormat := helper.FormatUser(user, token)
	response := helper.APIResponse("success", "User has been successfully added", http.StatusOK, userFormat)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input dto.LoginDTO
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("error", "Please Check Your Data", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error", "Login Failed, please check your email or password", http.StatusUnauthorized, errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error", "Login Failed, Failed to generate JWT Token", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helper.FormatUser(loggedUser, token)
	c.JSON(http.StatusOK, formatter)
}

func (h *userHandler) CheckEmailAvailable(c *gin.Context) {
	var input dto.EmailCheckDTO
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("error", "Please Check Your Data", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	available, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error", "Email checking failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if available {
		data := gin.H{"is_available": true}
		response := helper.APIResponse("success", "Email can be used", http.StatusOK, data)
		c.JSON(http.StatusOK, response)
		return
	}

	data := gin.H{"is_available": false}
	response := helper.APIResponse("success", "Email has been regestered", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	var response helper.Response
	file, err := c.FormFile("avatar")
	if err != nil {
		response = helper.APIResponse("error", "Failed to upload avatar", http.StatusBadRequest, gin.H{"is_uploaded": false})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := "images/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		response = helper.APIResponse("error", "Failed to upload avatar", http.StatusBadRequest, gin.H{"is_uploaded": false})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(entity.User)
	userID := currentUser.ID
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		response = helper.APIResponse("error", "Failed to update user", http.StatusBadRequest, gin.H{"is_uploaded": false})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = helper.APIResponse("success", "Avatar uploaded", http.StatusOK, gin.H{"is_uploaded": true})
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUserProfile(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(entity.User)
	userFormat := helper.FormatProfile(currentUser)
	response := helper.APIResponse("success", "Success Get User Data", http.StatusOK, userFormat)
	c.JSON(http.StatusOK, response)
}
