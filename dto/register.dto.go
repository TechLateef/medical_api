package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" `
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
}

type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password,omitempty" form:"password,omitempty"  binding:"required"`
}
