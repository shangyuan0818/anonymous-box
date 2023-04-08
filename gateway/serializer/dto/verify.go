package dto

type ApplyEmailVerify struct {
	Email string `json:"email" binding:"required,email"`
}
