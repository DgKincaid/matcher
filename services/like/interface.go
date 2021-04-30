package like

import (
	"matcher/entity"

	"github.com/google/uuid"
)

type Reader interface {
	List(userId uuid.UUID, offset int, limit int) ([]*entity.Like, error)
}

type Writer interface {
	Create(e *entity.Like) (uuid.UUID, error)
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	// GetLike(id uuid.UUID) (*entity.Like, error)
	CreateLike(fromId uuid.UUID, toId uuid.UUID) error
	ListLikes(userId uuid.UUID, page int, pageSize int) ([]*entity.Like, error)
}
