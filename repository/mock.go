package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupTestDB sets up a test database instance.
func SetupTestDB() *gorm.DB {
	connection := "user=postgres host=localhost password=password dbname=testdb sslmode=disable port=5432"
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database: " + err.Error())
	}
	// Auto-migrate tables for testing
	CreateTables(db) // This function creates tables if they don't exist
	return db
}
