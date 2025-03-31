package dto

type UserRegisterRequest struct {
	NIM         string `json:"nim" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Username    string `json:"username" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshATRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
