package app

import (
	"fmt"
	"log"
	"matcher/config"
	"matcher/db"

	"github.com/gin-gonic/gin"
)

// ServiceFactory
type ServiceFactory struct {
}

// RepositoryFactory stores repos
type RepositoryFactory struct {
}

// App struct
type App struct {
	DB         *db.DB
	router     *gin.Engine
	config     *config.Config
	Services   ServiceFactory
	Repository RepositoryFactory
}

// Initialize App
func Initialize() (*App, error) {
	config := config.Get()

	fmt.Println(config)

	db, err := db.Get(config.GetDBConnStr())

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &App{
		DB:     db,
		config: config,
	}, nil
}
