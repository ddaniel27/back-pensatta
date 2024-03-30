package postgres

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Name     string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
	Password string `env:"DB_PASSWORD"`
}

type PoolConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func NewGormPostgresClient(config Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(getConnectionString(config)), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return db
}

func getConnectionString(config Config) string {
	return fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s password=%s sslmode=disable",
		config.Host,
		config.User,
		config.Name,
		config.Port,
		config.Password,
	)
}
