package main

import (
	"log"
	"matcher/api/app"
	"matcher/api/handler"
	"matcher/repository"
	"matcher/services/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	a, err := app.Initialize()

	if err != nil {
		log.Fatal("Error loading application")
	}

	r := gin.Default()

	r.Use(cors.Default())

	userRepo := repository.NewUserPostgres(a.DB.Client)
	userService := user.NewUserService(userRepo)

	handler.UsersHandler(r, userService)

	r.Run(":3001")

	log.Println("Running....")
}
