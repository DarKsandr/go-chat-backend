package pkg

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
	dbErr      error
)

func InitDB() (*gorm.DB, error) {
	dbOnce.Do(func() {
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, database)
		dbInstance, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
			// Logger: logger.Default.LogMode(logger.Info),
		})
	})

	if dbErr != nil {
		return nil, dbErr
	}

	return dbInstance, nil
}

func OpenDB() *gorm.DB {
	db, err := InitDB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	return db
}
