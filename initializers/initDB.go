package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// getDSN retrieves the database connection string from environment variables.
func getDSN() string {
	return os.Getenv("DB_SECRET")
}

// ConnectToDb initializes the database connection.
func ConnectToDb() error {
	var err error

	dsn := getDSN()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	log.Println("Database connection established")
	return nil
}

//postgresql://aaaaaaaaa:gj4Ux7s2b77f9xWpxx8R2HNJ5jMdR24w@dpg-cqi8ou6ehbks73bprbp0-a.singapore-postgres.render.com/ab
