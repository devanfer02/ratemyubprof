package formatter

import (
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
)

func FormatUserEntityToDto(user *entity.User) dto.FetchUserResponse {
	return dto.FetchUserResponse{
		ID:        user.ID,
		NIM:       user.NIM,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
	}
}
func FormatToUserProfile(user *entity.User, reviews []entity.ReviewWithRelations) (dto.UserProfileResponse) {
	userDto :=  FormatUserEntityToDto(user)
	reviewDto := FormatReviewEntitiesToDto(reviews)

	return dto.UserProfileResponse{
		UserProfile: userDto,
		RecentReviews: reviewDto,
		ReviewsCount: len(reviewDto),
	}
}