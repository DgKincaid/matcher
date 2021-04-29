package main

import (
	"flag"
	"fmt"
	"log"
	"matcher/config"
	"matcher/db"
	"matcher/seeders/seeds"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}

	flag.Parse()

	args := flag.Args()

	config := config.Get()

	fmt.Println(config)

	db, err := db.Get(config.GetDBConnStr())

	if err != nil {
		log.Fatal("unable to establish db connection", err)
	}

	if len(args) > 0 {
		switch args[0] {
		case "up":
			seeds.Execute(db.Client, args[1:]...)
		case "down":
			log.Println("seeders down")
		}
	}
}
