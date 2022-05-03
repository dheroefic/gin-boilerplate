package database

import (
	"fmt"
	"os"
	"strconv"
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
		helpers.Logger("DATABASE", "Cannot open connection to database", true)
	}

	return db, err
}

func InitConnection() *gorm.DB {
	sqlConnection, _ = OpenConnection()
	sqlDB, _ := sqlConnection.DB()
	helpers.Logger("DATABASE", "Initializing database connection...", false)
	maxConn, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_CONNECTION"))
	maxLifeTime, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_LIFETIME"))
	sqlDB.SetMaxIdleConns(maxConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Minute)
	helpers.Logger("DATABASE", "Database connection has been established", false)
	return sqlConnection
}

func GetSession() gorm.DB {
	return *sqlConnection
}
