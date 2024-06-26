package db

import (
	"database/sql"
	"fmt"
	"modules-app/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	maxIdleConns = 20
	maxOpenConns = 20
	maxLifeTime  = 10 * time.Minute
)

var db *sql.DB

func DataBase(c config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySQLDB, c.MySQLHost, c.MySQLPassword, c.MySQLPort, c.MySQLUser,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return &gorm.DB{}, fmt.Errorf("failed to return a connection from the pool: %w", err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(maxLifeTime)
	sqlDB.SetMaxOpenConns(maxOpenConns)

	fmt.Println("Connected!")

	return db, nil
}
