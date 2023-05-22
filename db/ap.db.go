package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @dev Create a PostgresQL instance
//
// @return *gorm.DB
func EstablishPostgresClient() *gorm.DB {
	// prepare dsn
	dsn := os.Getenv("POSTGRESDB_DSN")
	if dsn == "" {
		log.Fatal("!POSTGRESDB_DSN - dsn is not defined.")
	}

	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("!PostgresQL Connection - Cannot conenct to PostgresQL server")
	}

	// return db
	log.Println("PostgresQL connected...")
	return db
}