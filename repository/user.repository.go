package repository

import (
	"log"
	model "medical_api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.Patient) model.Patient
	VerifyUser(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	UpdateProfile(user model.Patient, userId uint64) model.Patient
	FindByID(userID string) (model.Patient, error)
	DeleteUser(user model.Patient, userId uint64) model.Patient
	Save(user model.Patient) error
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(UserRepo *gorm.DB) UserRepository {
	return &userRepository{
		connection: UserRepo,
	}
}

func (db *userRepository) CreateUser(user model.Patient) model.Patient {
	db.connection.Save(&user)
	db.connection.Preload("Users").Find(&user)
	return user
}

func (db *userRepository) VerifyUser(email string, passsword string) interface{} {
	var user model.Patient
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user model.Patient
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userRepository) FindByID(userID string) (model.Patient, error) {
	var user model.Patient
	if err := db.connection.First(&user, "id = ?", userID).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (db *userRepository) Save(user model.Patient) error {

	return db.connection.Save(&user).Error
}

func (db *userRepository) UpdateProfile(_user model.Patient, userId uint64) model.Patient {
	var user model.Patient
	// Retrieve the user by ID first
	result := db.connection.First(&user, userId)
	if result.Error != nil {
		log.Fatalf("Failed to find user: %v", result.Error)
	}

	// Update the user fields with the new data, excluding the password
	user.Name = _user.Name
	user.Email = _user.Email
	user.Phone = _user.Phone

	// Save the updated user back to the database
	db.connection.Save(&user)

	// Load related data if necessary
	db.connection.Preload("Users").First(&user, userId)

	return user
}

func (db *userRepository) DeleteUser(_user model.Patient, userId uint64) model.Patient {
	var user model.Patient
	// Retrieve the user bt ID first
	result := db.connection.First(&user, userId)
	if result.Error != nil {
		log.Fatalf("Failed to find user: %v", result.Error)
	}

	// Delete the user from the database
	db.connection.Delete(&user)

	// Return the deleted user object
	return user

}
