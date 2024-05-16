package model

import "time"

type User struct {
	Id        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Phone     string    `gorm:"->;<-;not null" json:"-"`
	Password  string    `json:"password" binding:"required,min=6"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	Role      string    `gorm:"type:varchar(255)" json:"role"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
