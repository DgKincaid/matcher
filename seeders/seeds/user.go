package seeds

import (
	"fmt"
	"log"

	"github.com/bxcodec/faker/v3"

	"matcher/repository"
	"matcher/services/user"
)

func (s Seed) UserSeed() {
	userRepo := repository.NewUserPostgres(s.db)
	userService := user.NewUserService(userRepo)

	for i := 0; i < 100; i++ {
		firstName := faker.FirstName()
		lastName := faker.LastName()
		email := fmt.Sprintf("%s.%s@test.test", firstName, lastName)

		id, err := userService.CreateUser(firstName, lastName, email, faker.Password())

		if err != nil {
			log.Fatal("Error Seeding user", id, err)
		}

		log.Printf("Successfully created, %s", id)
	}
}
