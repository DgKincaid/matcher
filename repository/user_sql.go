package repository

import (
	"log"
	"matcher/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	db.AutoMigrate(&entity.User{})
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) Create(e *entity.User) (uuid.UUID, error) {
	result := r.db.Create(&e)

	if result.Error != nil {
		log.Printf("Error user create, %v", result.Error)
		return e.ID, result.Error
	}

	return e.ID, nil
}

func (r *UserPostgres) Get(id uuid.UUID) (*entity.User, error) {
	return &entity.User{}, nil
}
