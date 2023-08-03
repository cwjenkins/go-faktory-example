package config

import (
	"fmt"
	"time"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB;

func Init() {
	var err interface{}
	dbUser := "YOUR_USER"
	dbName := "YOUR_DB"

	dsn := fmt.Sprintf("host=localhost user=%s dbname=%s", dbUser, dbName)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	if err != nil {
		fmt.Println(err)
	}
}

func DB() *gorm.DB {
	return db;
}
