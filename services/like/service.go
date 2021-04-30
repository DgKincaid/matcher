package like

import (
	"matcher/entity"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewLikeService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateLike(fromId uuid.UUID, toId uuid.UUID) error {

	l, err := entity.NewLike(fromId, toId)

	if err != nil {
		return err
	}

	_, err = s.repo.Create(l)

	if err != nil {
		return err
	}

	return nil
}

// ListLikes Get all users who like you
func (s *Service) ListLikes(userId uuid.UUID, page int, pageSize int) ([]*entity.Like, error) {

	// Page size cant go over 100 as the max
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	likes, err := s.repo.List(userId, offset, pageSize)

	if err != nil {
		return nil, err
	}

	return likes, nil
}
