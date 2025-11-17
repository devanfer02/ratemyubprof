package dto

type UserRegisterRequest struct {
	NIM         string `json:"nim" validate:"required,max=16"`
	Password    string `json:"password" validate:"required"`
	Username    string `json:"username" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,alphanum,min=6,max=25"`
}

type ForgotPasswordRequest struct {
	NIM         string `json:"nim" validate:"required,max=16"`
	Password    string `json:"password" validate:"required"`
	Username    string `json:"username" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,alphanum,min=6,max=25"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserProfileRequest struct {
	UserID string
}

type UserProfileResponse struct {
	ID            string                `json:"id"`
	NIM           string                `json:"nim"`
	Username      string                `json:"username"`
	CreatedAt     string                `json:"createdAt"`
	ReviewsCount  int                   `json:"reviewsCount"`
	RecentReviews []FetchReviewResponse `json:"recentReviews"`
}

type FetchUserParams struct {
	Username string
	NIM      string
}

type FetchUserResponse struct {
	ID        string `json:"id"`
	NIM       string `json:"nim"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

type UserTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshATRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
