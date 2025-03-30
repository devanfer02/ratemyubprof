package dto

type UserRegisterRequest struct {
	NIM string `json:"nim"`
	Password string `json:"password"`
	Username string `json:"username"`
	NewPassword string `json:"newPassword"`
}