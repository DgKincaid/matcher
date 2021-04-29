package user

import (
	"log"
	"matcher/entity"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repository
}

func NewUserService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateUser(firstName string, lastName string, email string, pass string) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	u, err := entity.NewUser(firstName, lastName, email, string(hash))

	if err != nil {
		return u.ID, err
	}

	return u.ID, nil
}

func (s *Service) GetUser(id uuid.UUID) (*entity.User, error) {
	return s.repo.Get(id)
}
