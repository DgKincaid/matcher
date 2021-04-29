package seeds

import (
	"database/sql"
	"log"
	"reflect"
)

type Seed struct {
	db *sql.DB
}

func Execute(db *sql.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	if len(seedMethodNames) == 0 {
		log.Println("Running all seeders...")

		for i := 0; i < seedType.NumMethod(); i++ {
			method := seedType.Method(i)
			seed(s, method.Name)
		}
	}
}

func seed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)

	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}

	log.Println("Seeding ", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seeding ", seedMethodName, "succeded")
}
