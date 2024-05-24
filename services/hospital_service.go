package services

// import (
// 	"medical_api/dto"
// 	"medical_api/model"
// 	"medical_api/repository"
// )

// type HospitalService interface {
// 	GetAllHospitals() []dto.HospitalDTO
// 	GetHospitalByID(id uint64) (dto.HospitalDTO, error)
// 	GetDoctorsByHospitalID(id uint64) ([]model.Doctor, error)
// }

// type hospitalService struct {
// 	hospitalRepository repository.HospitalRepository
// }

// func NewHospitalService(repo repository.HospitalRepository) HospitalService {
// 	return &hospitalService{
// 		hospitalRepository: repo,
// 	}
// }

// func (s *hospitalService) GetAllHospitals() []dto.HospitalDTO {
// 	hospitals := s.hospitalRepository.GetAllHospitals()
// 	hospitalDTOs := make([]dto.HospitalDTO, len(hospitals))

// 	for i, hospital := range hospitals {
// 		hospitalDTOs[i] = s.convertToHospitalDTO(hospital)
// 	}
// 	return hospitalDTOs
// }

// func (s *hospitalService) GetHospitalByID(id uint64) (dto.HospitalDTO, error) {
// 	hospital, err := s.hospitalRepository.GetHospitalByID(id)
// 	if err != nil {
// 		return dto.HospitalDTO{}, err
// 	}
// 	return s.convertToHospitalDTO(hospital), nil
// }

// func (s *hospitalService) GetDoctorsByHospitalID(id uint64) ([]dto.DoctorDTO, error) {
// 	doctors, err := s.hospitalRepository.GetDoctorsByHospitalID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	doctorDTOs := make([]dto.DoctorDTO, len(doctors))
// 	for i, doctor := range doctors {
// 		doctorDTOs[i] = dto.DoctorDTO{
// 			ID:          doctor.ID,
// 			Name:        doctor.Name,
// 			Specialty:   doctor.Specialty,
// 			PhoneNumber: doctor.PhoneNumber,
// 			HospitalID:  doctor.HospitalID,
// 		}
// 	}
// 	return doctorDTOs, nil
// }

// func (s *hospitalService) convertToHospitalDTO(hospital model.Hospital) dto.HospitalDTO {
// 	doctorDTOs := make([]dto.DoctorDTO, len(hospital.Doctors))
// 	for i, doctor := range hospital.Doctors {
// 		doctorDTOs[i] = dto.DoctorDTO{
// 			ID:          doctor.ID,
// 			Name:        doctor.Name,
// 			Specialty:   doctor.Specialty,
// 			PhoneNumber: doctor.PhoneNumber,
// 			HospitalID:  doctor.HospitalID,
// 		}
// 	}

// 	return dto.HospitalDTO{
// 		ID:        hospital.ID,
// 		Name:      hospital.Name,
// 		Location:  hospital.Location,
// 		Doctors:   doctorDTOs,
// 		CreatedAt: hospital.CreatedAt.Format("2006-01-02 15:04:05"),
// 		UpdatedAt: hospital.UpdatedAt.Format("2006-01-02 15:04:05"),
// 	}
// }
