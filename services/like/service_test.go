package like

import (
	"matcher/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockRepo struct {
}

func (r *MockRepo) List(userId uuid.UUID, offset int, limit int) ([]*entity.Like, error) {
	return nil, nil
}

func (r *MockRepo) Create(e *entity.Like) (uuid.UUID, error) {
	return uuid.New(), nil
}

func NewMockRepo() Repository {
	return &MockRepo{}
}

func CreateLikeTest(t *testing.T) {
	fromId := uuid.New()
	toId := uuid.New()

	likeService := NewLikeService(NewMockRepo())

	err := likeService.CreateLike(fromId, toId)

	assert.NotNil(t, err)

}
