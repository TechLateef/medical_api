package repository

import (
	"log"
	"medical_api/model"

	"gorm.io/gorm"
)

type DoctorRepository interface {
	GetAllDoctors() []model.Doctor
	GetDoctorByID(doctorID uint64) model.Doctor
	CreateDoctor(doctor model.Doctor) model.Doctor
	UpdateDoctor(doctor model.Doctor, doctorId uint64) model.Doctor
	DeleteDoctor(doctorID uint64) model.Doctor
}

type doctorRepository struct {
	connection *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorRepository{
		connection: db,
	}
}

func (r *doctorRepository) GetAllDoctors() []model.Doctor {
	var doctors []model.Doctor
	r.connection.Find(&doctors)
	return doctors
}

func (r *doctorRepository) GetDoctorByID(doctorID uint64) model.Doctor {
	var doctor model.Doctor
	r.connection.First(&doctor, doctorID)
	return doctor
}

func (r *doctorRepository) CreateDoctor(doctor model.Doctor) model.Doctor {
	r.connection.Create(&doctor)
	return doctor
}

func (r *doctorRepository) UpdateDoctor(_doctor model.Doctor, doctorId uint64) model.Doctor {
	var doctor model.Doctor
	result := r.connection.First(&doctor, doctorId)
	if result.Error != nil {
		log.Fatalf("Failed to find user: %v", result.Error)
	}
	r.connection.Save(&doctor)

	// Load related data if necessary
	r.connection.Preload("Users").First(&doctor, doctorId)

	return doctor
}

func (r *doctorRepository) DeleteDoctor(doctorID uint64) model.Doctor {
	var doctor model.Doctor
	r.connection.Delete(&doctor, doctorID)
	return doctor
}
