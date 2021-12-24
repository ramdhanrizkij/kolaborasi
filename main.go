package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kolaborasi/config"
	"kolaborasi/handler"
	"kolaborasi/helper"
	"kolaborasi/repository"
	"kolaborasi/service"
	"net/http"
	"strings"
)

var (
	db              *gorm.DB                      = config.SetupDatabaseConnection()
	userRepo        repository.UserRepository     = repository.NewUserRepository(db)
	userService     service.UserService           = service.NewUserService(userRepo)
	authService     service.AuthService           = service.NewAuthService()
	userHandler                                   = handler.NewUserHandler(userService, authService)
	campaignRepo    repository.CampaignRepository = repository.NewCampaignRepository(db)
	campaignService service.CampaignService       = service.NewCampaignService(campaignRepo)
	campaignHandler                               = handler.NewCampaignHandler(campaignService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.Static("/imagse", "./images")

	authAPI := r.Group("v1/auth")
	authAPI.POST("/register", userHandler.RegisterUser)
	authAPI.POST("/login", userHandler.LoginUser)
	authAPI.POST("/email_checkers", userHandler.CheckEmailAvailable)
	authAPI.GET("/me", authMiddleware(authService, userService), userHandler.GetUserProfile)

	userAPI := r.Group("v1/user")
	userAPI.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

	campaignAPI := r.Group("v1/campaign")
	campaignAPI.GET("/", campaignHandler.GetCampaignData)
	campaignAPI.GET("/:id", campaignHandler.GetCampaignDetail)
	campaignAPI.POST("/", authMiddleware(authService, userService), campaignHandler.CreateCampaign)

	r.Run(":8000")
}

func authMiddleware(authService service.AuthService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("error", "Unautorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("error", "Unautorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("error", "Unautorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userID)
		if err != nil {
			response := helper.APIResponse("error", "Unautorized", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}

}
