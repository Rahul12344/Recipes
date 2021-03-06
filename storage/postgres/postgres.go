package postgres

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// postgres gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Table interface for creating new tables
type Table interface {
	create()
}

//CreatePostgresTables creates postgresql tables
func CreatePostgresTables(tables ...Table) {
	for _, table := range tables {
		table.create()
	}
}

// Connect connects to postgresql; if hosted option is flagged, connects (for now) to Heroku table
func Connect(hosted bool, dbEnv string, databaseHost string, username string, databaseName string, password string) *gorm.DB {
	if hosted {
		db, err := gorm.Open("postgres", os.Getenv(dbEnv))
		if err != nil {
			print(err.Error())
		}
		//defer db.Close()
		fmt.Println("Successfully connected!", db)
		return db
	}
	db, err := gorm.Open("postgres", dbURI(databaseHost, username, databaseName, password))
	if err != nil {
		print(err.Error())
	}
	//defer db.Close()
	fmt.Println("Successfully connected!", db)
	return db
}

func dbURI(databaseHost string, username string, databaseName string, password string) string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)
}
