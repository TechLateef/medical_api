package router

import (
	"medical_api/config"
	"medical_api/controllers"
	"medical_api/repository"
	"medical_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepo repository.UserRepository = repository.NewUserRepository(db)

	userServ   services.AuthService       = services.NewAuthService(userRepo)
	jwtServ    services.JWTService        = services.NewJWTService()
	controller controllers.AuthController = controllers.NewAuthController(userServ, jwtServ)
)

func Routes() {
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Application is up and running")
	})

	authRoute := route.Group("/api/v1/auth")
	{
		authRoute.POST("/create", controller.Register)
		authRoute.POST("/login", controller.Login)
		authRoute.POST("/logout", controller.Logout)
	}
	// Run route whenever triggered
	route.Run()
}
