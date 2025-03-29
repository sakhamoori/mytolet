package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID        string  `gorm:"type:uuid;primary_key" json:"id"`
	Street    string  `gorm:"not null" json:"street"`
	City      string  `gorm:"not null" json:"city"`
	State     string  `gorm:"not null" json:"state"`
	ZipCode   string  `gorm:"not null" json:"zipCode"`
	Country   string  `gorm:"not null" json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (a *Address) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return nil
}