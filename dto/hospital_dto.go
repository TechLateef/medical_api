package dto

type CreateHospitalDTO struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

type UpdateHospitalDTO struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type HospitalDTO struct {
	ID        uint64      `json:"id"`
	Name      string      `json:"name"`
	Location  string      `json:"location"`
	Doctors   []DoctorDTO `json:"doctors"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
}
