package repository

import (
	"medical_api/model"

	"gorm.io/gorm"
)

type HospitalRepository interface {
	GetAllHospitals() []model.Hospital
	GetHospitalByID(id uint64) (model.Hospital, error)
	GetDoctorsByHospitalID(id uint64) ([]model.Doctor, error)
}

type hospitalRepository struct {
	connection *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) HospitalRepository {
	return &hospitalRepository{
		connection: db,
	}
}

func (r *hospitalRepository) GetAllHospitals() []model.Hospital {
	var hospitals []model.Hospital
	r.connection.Preload("Doctors").Find(&hospitals)
	return hospitals
}

func (r *hospitalRepository) GetHospitalByID(id uint64) (model.Hospital, error) {
	var hospital model.Hospital
	err := r.connection.Preload("Doctors").First(&hospital, id).Error
	if err != nil {
		return model.Hospital{}, err
	}
	return hospital, nil
}
func (r *hospitalRepository) GetDoctorsByHospitalID(id uint64) ([]model.Doctor, error) {
	var doctors []model.Doctor
	err := r.connection.Where("hospital_id = ?", id).Find(&doctors).Error
	if err != nil {
		return nil, err
	}
	return doctors, nil
}
