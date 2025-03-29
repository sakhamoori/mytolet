package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PropertyStatus string

const (
	PropertyStatusAvailable PropertyStatus = "AVAILABLE"
	PropertyStatusRented    PropertyStatus = "RENTED"
	PropertyStatusPending   PropertyStatus = "PENDING"
	PropertyStatusInactive  PropertyStatus = "INACTIVE"
)

type Property struct {
	ID            string         `gorm:"type:uuid;primary_key" json:"id"`
	Title         string         `gorm:"not null" json:"title"`
	Description   string         `gorm:"not null" json:"description"`
	AddressID     string         `gorm:"type:uuid" json:"-"`
	Address       Address        `gorm:"foreignKey:AddressID" json:"address"`
	Bedrooms      int            `json:"bedrooms"`
	Bathrooms     float64        `json:"bathrooms"`
	Price         float64        `gorm:"not null" json:"price"`
	AvailableFrom time.Time      `json:"availableFrom"`
	Amenities     []string       `gorm:"type:text[]" json:"amenities"`
	Images        []string       `gorm:"type:text[]" json:"images"`
	OwnerID       string         `gorm:"type:uuid" json:"-"`
	Owner         User           `gorm:"foreignKey:OwnerID" json:"owner"`
	Status        PropertyStatus `gorm:"not null" json:"status"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}

func (p *Property) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	if p.Status == "" {
		p.Status = PropertyStatusAvailable
	}
	return nil
}