package model

import (
	"time"
)

type Doctor struct {
	ID             uint64          `gorm:"primary_key:auto_increment" json:"id"`
	Name           string          `gorm:"type:varchar(255)" json:"name"`
	MedicalRecords []MedicalRecord `gorm:"foreignKey:DoctorID" json:"medical_records"`
	Specialty      string          `gorm:"type:varchar(255)" json:"specialty"`
	CreatedAt      time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}
