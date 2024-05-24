package dto

type CreateDoctorDTO struct {
	Id          uint64 `uri:"id" form:"id" binding:"required,uuid"`
	Name        string `json:"name" binding:"required"`
	Specialty   string `json:"specialty" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UpdateDoctorDTO struct {
	Id          uint64 `uri:"id" form:"id" binding:"required,uuid"`
	Name        string `json:"name"`
	Specialty   string `json:"specialty"`
	PhoneNumber string `json:"phone_number"`
}

type DoctorDTO struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Specialty   string `json:"specialty"`
	PhoneNumber string `json:"phone_number"`
	HospitalID  uint64 `json:"hospital_id"`
}
