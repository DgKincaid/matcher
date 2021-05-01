package app

import (
	"fmt"
	"log"
	"matcher/config"
	"matcher/db"
)

// App struct used to create a default application
type App struct {
	DB     *db.DB
	config *config.Config
}

// Initialize App connects to the database
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
