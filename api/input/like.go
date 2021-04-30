package input

import "github.com/google/uuid"

type CreateLike struct {
	FromID uuid.UUID `json:"fromId" binding:"required"`
	ToID   uuid.UUID `json:"toId" binding:"required"`
}

type ListLikes struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}
