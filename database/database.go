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

var sqlConnection *gorm.DB

func InitConnection() *gorm.DB {
	var err error
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)

	helpers.Logger("DATABASE", "Initializing database connection...", false)

	sqlConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		helpers.Logger("DATABASE", "Cannot open connection to database", true)
	}

	sqlDB, _ := sqlConnection.DB()
	maxConn, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_CONNECTION"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNECTION"))
	maxLifeTime, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_LIFETIME"))
	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Minute)

	helpers.Logger("DATABASE", "Database connection has been established", false)
	return sqlConnection
}

func GetSession() *gorm.DB {
	return sqlConnection.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})
}
