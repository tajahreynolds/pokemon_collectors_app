package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global variable to store the database connection

type User struct {
	gorm.Model
	Name 	  string
	Email 	  string
}


// https://github.com/go-gorm/postgres
func ConnectPostgres(dsn string) error {
    var err error
    DB, err = gorm.Open(postgres.New(postgres.Config{
        DSN: dsn,
        PreferSimpleProtocol: true, // disables implicit prepared statement usage
    }), &gorm.Config{})

    if err != nil {
        return err
    }

    // Migrate the schema
    DB.AutoMigrate(&User{})

    return nil
}

// FindUser fetches a user from the database by email
func FindUser(email string) *User {
    var user User
    if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
        return nil
    }
    return &user
}

// CreateUser adds a new user to the database
func CreateUser(user User) error {
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}