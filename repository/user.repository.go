package repository

import (
	model "medical_api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) model.User
	VerifyUser(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByID(userID string) (model.User, error)
	Save(user model.User) error
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(UserRepo *gorm.DB) UserRepository {
	return &userRepository{
		connection: UserRepo,
	}
}

func (db *userRepository) CreateUser(user model.User) model.User {
	db.connection.Save(&user)
	db.connection.Preload("Users").Find(&user)
	return user
}

func (db *userRepository) VerifyUser(email string, passsword string) interface{} {
	var user model.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userRepository) FindByID(userID string) (model.User, error) {
	var user model.User
	if err := db.connection.First(&user, "id = ?", userID).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (db *userRepository) Save(user model.User) error {

	return db.connection.Save(&user).Error
}
