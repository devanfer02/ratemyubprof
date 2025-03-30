package dto

type UserRegisterRequest struct {
	NIM         string `json:"nim" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Username    string `json:"username" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
