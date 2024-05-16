package services

import (
	"log"
	"medical_api/dto"
	"medical_api/model"
	"medical_api/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) model.User
	IsDuplicateEmail(email string) bool
	ClearUserToken(userID string) error
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (service *authService) CreateUser(user dto.RegisterDTO) model.User {
	userToCreate := model.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.CreateUser(userToCreate)
	return res
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyUser(email, password)
	if v, ok := res.(model.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false

}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *authService) ClearUserToken(userID string) error {
	user, err := service.userRepository.FindByID(userID)
	if err != nil {
		return err
	}

	user.Token = ""

	return service.userRepository.Save(user)
}
