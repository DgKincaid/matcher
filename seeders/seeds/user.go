package seeds

import (
	"fmt"
	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"

	"matcher/entity"
	"matcher/repository"
	"matcher/services/like"
	"matcher/services/user"
)

func DefaultUser() *entity.User {

	u := &entity.User{
		ID:        uuid.MustParse("fb8fb77f-295d-4126-bef9-9aea2b93ae43"),
		FirstName: "Tod",
		LastName:  "Ferry",
		Email:     "Tod.Ferry@test.test",
		Pass:      "$2a$04$htOXfCUldHvcW7cgJamiz.HqkW5nL/IcnhsrV/QRqTxH4m/NVhTFW",
	}

	return u
}

func (s Seed) UserSeed() {
	userRepo := repository.NewUserPostgres(s.db)
	userService := user.NewUserService(userRepo)

	likeRepo := repository.NewLikePostgres(s.db)
	likeService := like.NewLikeService(likeRepo)

	u := DefaultUser()

	// userRepo.Create(u)

	for i := 0; i < 100; i++ {
		firstName := faker.FirstName()
		lastName := faker.LastName()

		email := fmt.Sprintf("%s.%s@test.test", firstName, lastName)

		id, err := userService.CreateUser(firstName, lastName, email, faker.Password())

		likeService.CreateLike(id, u.ID)

		if err != nil {
			log.Fatal("Error Seeding user", id, err)
		}

		log.Printf("Successfully created, %s", id)
	}
}
