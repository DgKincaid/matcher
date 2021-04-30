package seeds

import (
	"log"
	"matcher/entity"
	"reflect"

	"gorm.io/gorm"
)

type Seed struct {
	// TODO: Update to use internal DB struct
	db *gorm.DB
}

func Execute(db *gorm.DB, seedMethodNames ...string) {
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

// Teardown drops all tables
func Teardown(db *gorm.DB) {
	if db.Migrator().HasTable(&entity.User{}) {
		db.Migrator().DropTable(&entity.User{})
	}

	if db.Migrator().HasTable(&entity.Like{}) {
		db.Migrator().DropTable(&entity.Like{})
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
