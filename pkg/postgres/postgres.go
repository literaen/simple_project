package postgres

import (
	"fmt"

	"github.com/literaen/simple_project/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GDB struct {
	DB *gorm.DB
}

func NewGDB(creds *config.DB_CREDS) (*GDB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		creds.Host,
		creds.User,
		creds.Password,
		creds.Name,
		creds.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return &GDB{DB: db}, nil
}
