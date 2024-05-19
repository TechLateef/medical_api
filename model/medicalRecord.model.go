package model

import (
	"time"
)

type MedicalRecord struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	PatientID uint64    `gorm:"not null" json:"patient_id"`
	Patient   Patient   `gorm:"foreignKey:PatientID" json:"patient"`
	Diagnosis string    `gorm:"type:varchar(255)" json:"diagnosis"`
	Treatment string    `gorm:"type:varchar(255)" json:"treatment"`
	DoctorID  uint64    `gorm:"not null" json:"doctor_id"`
	Doctor    Doctor    `gorm:"foreignKey:DoctorID" json:"doctor"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
