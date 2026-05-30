package database

import (
	"fmt"

	"github.com/kyawzinkoko-dev/online-learning-platform/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabae(config *configs.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
		config.DBSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
