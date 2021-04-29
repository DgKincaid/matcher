package repository

import (
	"database/sql"
	"matcher/entity"

	"github.com/google/uuid"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) Create(e *entity.User) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (r *UserPostgres) Get(id uuid.UUID) (*entity.User, error) {
	return &entity.User{}, nil
}
