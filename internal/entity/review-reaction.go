package entity

type ReactionType uint

const (
	LikeReactionType    ReactionType = 1
	DislikeReactionType ReactionType = 2
)

func ToReactionType(reaction string) ReactionType {
	switch reaction {
	case "like":
		return LikeReactionType
	case "dislike":
		return DislikeReactionType
	default:
		return 0
	}
}

type ReviewReaction struct {
	UserID    string `db:"user_id"`
	ReviewID  string `db:"review_id"`
	Type      ReactionType   `db:"reaction_type"`
	CreatedAt string `db:"created_at"`
}
