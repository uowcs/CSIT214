package airline

import (
	"github.com/uowcs/CSIT214/internal/models"
	"gorm.io/gorm"
)

type AirlineStorage struct {
	db *gorm.DB
}

func NewAirlineStorage(db *gorm.DB) *AirlineStorage {
	return &AirlineStorage{db: db}
}

func (a *AirlineStorage) CreateSchemas() error {
	// Create the table
	err := a.db.AutoMigrate(&models.Product{})
	if err != nil {
		return err
	}
	return nil
}
