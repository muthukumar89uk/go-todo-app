package driver

import (
	//user defined package
	"todo/helper"

	//built in package
	"fmt"
	"os"

	//third party package
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connection with postgres database
func DatabaseConnection() *gorm.DB {
	err := helper.Configure(".env")
	if err != nil {
		fmt.Println("error is loading env file ")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")

	//connecting to postgres-SQL
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		fmt.Println("error in connecting with database")
	}
	fmt.Printf("database connection successfull\n")
	return Database
}
