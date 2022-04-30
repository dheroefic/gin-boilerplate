package database

import (
	"fmt"
	"os"
	"time"

	"github.com/dheroefic/gin-boilerplate/utils/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	DBInterface interface {
		OpenConnection() (*gorm.DB, error)
	}
)

var sqlConnection *gorm.DB

func OpenConnection() (*gorm.DB, error) {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		helpers.Logger("DATABASE OPEN CONNECTION", "Cannot open connection to database", true)
	}

	return db, err
}

func InitConnection() *gorm.DB {
	sqlConnection, _ = OpenConnection()
	sqlDB, _ := sqlConnection.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	return sqlConnection
}

func GetSession() gorm.DB {
	return *sqlConnection
}
