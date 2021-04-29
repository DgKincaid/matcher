package user

import (
	"matcher/entity"

	"github.com/google/uuid"
)

type Reader interface {
	Get(id uuid.UUID) (*entity.User, error)
}

type Writer interface {
	Create(e *entity.User) (uuid.UUID, error)
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	GetUser(id uuid.UUID) (*entity.User, error)
	CreateUser(firstName string, lastName string, email string, pass string) (uuid.UUID, error)
}
