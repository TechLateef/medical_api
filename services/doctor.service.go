// services/doctor_service_impl.go
package services

import (
	"medical_api/dto"
	"medical_api/model"
	"medical_api/repository"
)

type DoctorService interface {
	GetDoctorByID(doctorID uint64) model.Doctor
	CreateDoctor(doctor dto.CreateDoctorDTO) model.Doctor
	UpdateDoctor(doctor dto.UpdateDoctorDTO, doctorId uint64) model.Doctor
	GetAllDoctors() []model.Doctor
	DeleteDoctor(doctorID uint64) model.Doctor
}

type doctorService struct {
	doctorRepository repository.DoctorRepository
}

func NewDoctorService(repo repository.DoctorRepository) DoctorService {
	return &doctorService{
		doctorRepository: repo,
	}
}

func (s *doctorService) GetAllDoctors() []model.Doctor {
	return s.doctorRepository.GetAllDoctors()
}

func (s *doctorService) GetDoctorByID(doctorID uint64) model.Doctor {
	return s.doctorRepository.GetDoctorByID(doctorID)
}

func (s *doctorService) CreateDoctor(doctor dto.CreateDoctorDTO) model.Doctor {
	doctorTocreate := model.Doctor{}
	return s.doctorRepository.CreateDoctor(doctorTocreate)
}

func (s *doctorService) UpdateDoctor(doctor dto.UpdateDoctorDTO, doctorId uint64) model.Doctor {
	updateDoctor := model.Doctor{}
	return s.doctorRepository.UpdateDoctor(updateDoctor, doctorId)
}

func (s *doctorService) DeleteDoctor(doctorID uint64) model.Doctor {
	return s.doctorRepository.DeleteDoctor(doctorID)
}
