// dto/medical_record_dto.go
package dto

type CreateMedicalRecordDTO struct {
	PatientID uint64 `json:"patient_id" binding:"required"`
	Diagnosis string `json:"diagnosis" binding:"required"`
	Treatment string `json:"treatment" binding:"required"`
	DoctorID  uint64 `json:"doctor_id" binding:"required"`
}

type UpdateMedicalRecordDTO struct {
	PatientID uint64 `json:"patient_id"`
	Diagnosis string `json:"diagnosis"`
	Treatment string `json:"treatment"`
	DoctorID  uint64 `json:"doctor_id"`
}
