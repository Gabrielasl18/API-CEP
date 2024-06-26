package storage

import (
	"fmt"
	"modules-app/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	maxIdleConns = 20
	maxOpenConns = 20
	maxLifetime  = 10 * time.Minute
)

func NewDB(c config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySQLUser, c.MySQLPassword, c.MySQLHost, c.MySQLPort, c.MySQLDB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to return a connection from the pool: %w", err)
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(maxLifetime)

	return db, nil
}
