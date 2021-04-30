package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Like struct {
	FromID    uuid.UUID `gorm:"type:uuid;primaryKey;"`
	ToID      uuid.UUID `gorm:"type:uuid;primaryKey;"`
	From      User      `gorm:"foreignKey:FromID;"`
	To        User      `gorm:"foreignKey:ToID;"`
	Status    string
	CreatedAt time.Time
}

func NewLike(fromId uuid.UUID, toId uuid.UUID) (*Like, error) {
	l := &Like{
		FromID:    fromId,
		ToID:      toId,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	return l, nil
}

func (l *Like) Validate() error {

	if l.FromID == l.ToID {
		return fmt.Errorf("invalid like")
	}

	return nil
}
