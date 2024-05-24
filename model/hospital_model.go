package model

import "time"

type Hospital struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Location  string    `gorm:"type:varchar(255)" json:"location"`
	Doctors   []Doctor  `gorm:"foreignKey:DoctorId" json:"doctors"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
