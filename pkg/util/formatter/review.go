package formatter

import (
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
)

func FormatReviewToEntity(review *dto.ProfessorReviewRequest) entity.Review {
	return entity.Review{
		UserID:       review.UserID,
		ProfessorID:  review.ProfessorID,
		Comment:      review.Comment,
		DiffRate:     review.DiffRate,
		FriendlyRate: review.FriendlyRate,
	}
}

func FormatReviewEntitiesToDto(reviews []entity.ReviewWithRelations) []dto.FetchReviewResponse {
	res := make([]dto.FetchReviewResponse, len(reviews))
	for idx, review := range reviews {
		res[idx] = FormatReviewEntityToDto(&review)
	}

	return res
}

func FormatReviewEntityToDto(review *entity.ReviewWithRelations) dto.FetchReviewResponse {
	return dto.FetchReviewResponse{
		ID:           review.ID,
		UserID:       review.UserID,
		ProfessorID:  review.ProfessorID,
		Comment:      review.Comment,
		DiffRate:     review.DiffRate,
		FriendlyRate: review.FriendlyRate,
		Like:         review.LikeCounter,
		Dislike:      review.DislikeCounter,
		CreatedAt:    review.CreatedAt.String(),
		IsLiked:      review.IsLiked,
		Professor:    FormatProfessorEntityToDto(review.Professor),
		User:         FormatUserEntityToDto(&review.User),
	}
}
