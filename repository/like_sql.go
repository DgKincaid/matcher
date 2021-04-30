package repository

import (
	"log"
	"matcher/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LikePostgres struct {
	db *gorm.DB
}

func NewLikePostgres(db *gorm.DB) *LikePostgres {
	db.AutoMigrate(&entity.Like{})
	return &LikePostgres{
		db: db,
	}
}

func (r *LikePostgres) Create(e *entity.Like) (uuid.UUID, error) {

	result := r.db.Create(&e)

	if result.Error != nil {
		log.Printf("Error like create, %v", result.Error)
		return e.FromID, result.Error
	}

	return e.FromID, nil
}

func (r *LikePostgres) List(userId uuid.UUID, offset int, limit int) ([]*entity.Like, error) {

	var likes []*entity.Like

	result := r.db.Model(&entity.Like{}).Offset(offset).Limit(limit).Where(
		"to_id = (?) AND STATUS = (?)", userId, "pending").Order(
		"created_at desc").Preload("From").Find(&likes)

	if result.Error != nil {
		log.Printf("Error like create, %v", result.Error)
		return nil, result.Error
	}

	return likes, nil
}
