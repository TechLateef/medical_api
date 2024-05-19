package controllers

import (
	"log"
	"medical_api/dto"
	"medical_api/model"
	"medical_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtServive  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtServive:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDTO
	errDto := ctx.ShouldBind(&loginDto)
	if errDto != nil {
		ctx.JSON(http.StatusBadRequest, errDto)
		return
	}

	authResult := c.authService.VerifyCredential(loginDto.Email, loginDto.Password)

	if user, ok := authResult.(model.Patient); ok {
		generateToken := c.jwtServive.GeneratedToken(user.Role, strconv.FormatUint(user.Id, 10))
		user.Token = generateToken
		ctx.JSON(http.StatusOK, user)
		return
	}

	// If VerifyCredential returns an error, handle it here
	if err, ok := authResult.(error); ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
}

func (c *authController) Logout(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	if tokenString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	userID, err := c.jwtServive.GetUserIDFromToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	err = c.authService.ClearUserToken(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDto dto.RegisterDTO
	errDto := ctx.ShouldBind(&registerDto)
	if errDto != nil {
		ctx.JSON(http.StatusBadRequest, errDto)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDto.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email already Exist "})

	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.MinCost)
		if err != nil {
			log.Fatalf("failed to hash password %v ", err)
		}
		registerDto.Password = string(hash)
		registerDto.Role = "patient"

		createUser := c.authService.CreateUser(registerDto)

		token := c.jwtServive.GeneratedToken(createUser.Role, strconv.FormatUint(createUser.Id, 10))
		createUser.Token = token
		ctx.JSON(http.StatusOK, createUser)

	}
}

func (c *authController) UpdateProfile(ctx *gin.Context) {
	var updateUser dto.UpdateUserDto
	ctx.ShouldBind(&updateUser)
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	updateUser.Id = id
	result := c.authService.UpdateProfile(updateUser, id)

	ctx.JSON(http.StatusBadRequest, result)
}

func (c *authController) DeleteUser(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var user model.Patient
	user.Id = id
	res := c.authService.DeleteUser(user, id)
	ctx.JSON(http.StatusBadRequest, res)

}
