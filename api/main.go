package main

import (
	"log"
	"matcher/api/app"
	"matcher/api/handler"
	"matcher/api/middleware"
	"matcher/repository"
	"matcher/services/like"
	"matcher/services/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main entrypoint for the api application. Responsible for initializing services/repos
func main() {
	// Using godotenv loads env variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}

	a, err := app.Initialize()

	if err != nil {
		log.Fatal("Error loading application")
	}

	r := gin.Default()

	r.Use(cors.Default())
	r.Use(middleware.Auth())

	userRepo := repository.NewUserPostgres(a.DB.Client)
	userService := user.NewUserService(userRepo)

	likeRepo := repository.NewLikePostgres(a.DB.Client)
	likeService := like.NewLikeService(likeRepo)

	handler.UsersHandler(r, userService)
	handler.LikesHandler(r, likeService)

	r.Run(":3001")
}
