package formatter

import (
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
)

func FormatReactionToEntity(reaction *dto.ReviewReactionRequest) *entity.ReviewReaction {
	return &entity.ReviewReaction{
		ReviewID: reaction.ReviewID,
		UserID:   reaction.UserID,
		Type:     entity.ToReactionType(reaction.Type),
	}
}