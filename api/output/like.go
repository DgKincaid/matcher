package output

import "time"

type ListLikes struct {
	Name      string    `json:"name" binding:"required"`
	FromID    string    `json:"fromId" binding:"required"`
	ToID      string    `json:"toId" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}
