package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"nurul-iman-blok-m/announcement"
	"nurul-iman-blok-m/auth"
	"nurul-iman-blok-m/database"
	"nurul-iman-blok-m/handler"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/role"
	"nurul-iman-blok-m/study_rundown"
	"nurul-iman-blok-m/user"
	"os"
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
	studyRundownHandler := handler.NewHandlerStudyRundown(studyRundownService)

	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup gin app
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")

	// setup s3 uploader
	cfg, errS3 := config.LoadDefaultConfig(context.TODO())
	if errS3 != nil {
		log.Printf("error: %v", errS3)
		return
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	announcementHandler := handler.NewHandlerAnnouncement(announcementService, *uploader, *configS3())

	api := router.Group("/api/v1")
	// for test api
	api.GET("/test", userHandler.RegisterUser)
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

	//roleInsert := model.Role{
	//	RoleName:  "super-admin",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//}
	//db.Save(&roleInsert)

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

func configS3() *s3.Client {

	creds := credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(os.Getenv("AWS_REGION")))

	if err != nil {
		log.Fatal(err)
	}

	return s3.NewFromConfig(cfg)

}
