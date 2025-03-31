package formatter

import (
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
)

func FormatReviewToEntity(review *dto.ProfessorReviewRequest) entity.Review {
	return entity.Review{
		UserID: review.UserID,
		ProfessorID: review.ProfessorID,
		Comment:      review.Comment,
		DiffRate: review.DiffRate,
		FriendlyRate: review.FriendlyRate,
	}
}