package output

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Email     string    `json:"email" binding:"required"`
}
