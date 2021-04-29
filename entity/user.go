package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User Struct
type User struct {
	gorm.Model
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Pass      string // hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(firstName string, lastName string, email string, pass string) (*User, error) {
	u := &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Pass:      pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.Validate()

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Validate() error {
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Pass == "" {
		return fmt.Errorf("invalid user entity.{ Email: %v }", u.Email)
	}

	return nil
}
