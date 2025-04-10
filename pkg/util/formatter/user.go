package formatter

import (
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
)

func FormatUserEntityToDto(user *entity.User) dto.FetchUserResponse {
	return dto.FetchUserResponse{
		ID:        user.ID,
		Username: user.Username,
		CreatedAt:  user.CreatedAt.String(),
	}
}