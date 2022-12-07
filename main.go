package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"nurul-iman-blok-m/announcement"
	"nurul-iman-blok-m/auth"
	"nurul-iman-blok-m/database"
	"nurul-iman-blok-m/handler"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/role"
	"nurul-iman-blok-m/study_rundown"
	"nurul-iman-blok-m/user"
	"strings"
)

func main() {
	db := database.Db()

	userRepository := user.NewRepository(db)
	roleRepository := role.NewRepository(db)
	announcementRepository := announcement.NewRepositoryAnnouncement(db)
	studyRundownRepository := study_rundown.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	roleService := role.NewRoleService(roleRepository)
	announcementService := announcement.NewServiceAnnouncement(announcementRepository)
	studyRundownService := study_rundown.NewService(studyRundownRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	roleHandler := handler.NewRoleHandler(roleService)
	announcementHandler := handler.NewHandlerAnnouncement(announcementService)
	studyRundownHandler := handler.NewHandlerStudyRundown(studyRundownService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")
	api.POST("/user/register", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.LoginUser)

	api.POST("/role/add", authMiddleware(authService, userService), roleHandler.SaveRole)
	api.GET("/roles", authMiddleware(authService, userService), roleHandler.GetRoles)

	api.POST("/announcement/add", authMiddleware(authService, userService), announcementHandler.AddAnnouncement)
	api.GET("/announcements", announcementHandler.GetAllAnnouncement)
	api.GET("/announcements/:id", announcementHandler.GetDetailAnnouncement)
	api.DELETE("/announcements/:id", authMiddleware(authService, userService), announcementHandler.DeleteAnnouncement)
	api.PUT("/announcements/:id", authMiddleware(authService, userService), announcementHandler.UpdateAnnouncement)

	api.GET("/user/ustadz", authMiddleware(authService, userService), studyRundownHandler.GetListUstadzName)
	api.POST("/rundown/add", authMiddleware(authService, userService), studyRundownHandler.AddStudy)
	api.GET("/rundown", studyRundownHandler.GetAllRundown)
	api.GET("/rundown/:id", studyRundownHandler.GetDetailStudyRundown)
	api.DELETE("/rundown/:id", authMiddleware(authService, userService), studyRundownHandler.DeleteStudyRundown)
	api.PUT("/rundown/:id", authMiddleware(authService, userService), studyRundownHandler.UpdateStudyRundown)

	router.Run(":8080")
}

func authMiddleware(autService auth.Service, userService user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := autService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := uint(claim["user_id"].(float64))

		currentUser, errFindUser := userService.GetUserByID(userId)

		if errFindUser != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", currentUser)
	}

}
