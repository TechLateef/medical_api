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

	// User Repository and Services
	userRepo repository.UserRepository  = repository.NewUserRepository(db)
	userServ services.AuthService       = services.NewAuthService(userRepo)
	jwtServ  services.JWTService        = services.NewJWTService()
	authCtrl controllers.AuthController = controllers.NewAuthController(userServ, jwtServ)

	// Doctor Repository and Services
	doctorRepo repository.DoctorRepository  = repository.NewDoctorRepository(db)
	doctorServ services.DoctorService       = services.NewDoctorService(doctorRepo)
	doctorCtrl controllers.DoctorController = controllers.NewDoctorController(doctorServ)

	// Medical Record Repository and Services
	medicalRecordRepo repository.MedicalRecordRepository  = repository.NewMedicalRecordRepository(db)
	medicalRecordServ services.MedicalRecordService       = services.NewMedicalRecordService(medicalRecordRepo)
	medicalRecordCtrl controllers.MedicalRecordController = controllers.NewMedicalRecordController(medicalRecordServ)
)

func Routes() {
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Application is up and running")
	})

	// Auth Routes
	authRoute := route.Group("/api/v1/auth")
	{
		authRoute.POST("/create", authCtrl.Register)
		authRoute.POST("/login", authCtrl.Login)
		authRoute.POST("/logout", authCtrl.Logout)
	}

	// Doctor Routes
	doctorRoute := route.Group("/api/v1/doctors")
	{
		doctorRoute.GET("/", doctorCtrl.GetAllDoctors)
		doctorRoute.POST("/", doctorCtrl.CreateDoctor)
		doctorRoute.GET("/:id", doctorCtrl.FindDoctorByID)
		doctorRoute.PUT("/:id", doctorCtrl.UpdateDoctor)
		doctorRoute.DELETE("/:id", doctorCtrl.DeleteDoctor)
	}

	// Medical Record Routes
	medicalRecordRoute := route.Group("/api/v1/medical-records")
	{
		medicalRecordRoute.GET("/", medicalRecordCtrl.GetAllMedicalRecords)
		medicalRecordRoute.POST("/", medicalRecordCtrl.CreateMedicalRecord)
		medicalRecordRoute.GET("/:id", medicalRecordCtrl.FindMedicalRecordByID)
		medicalRecordRoute.PATCH("/:id", medicalRecordCtrl.UpdateMedicalRecord)
		medicalRecordRoute.DELETE("/:id", medicalRecordCtrl.DeleteMedicalRecord)
	}

	// Run route whenever triggered
	route.Run()
}
