package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentType string

const (
	DocumentTypeIDProof        DocumentType = "ID_PROOF"
	DocumentTypeIncomeProof    DocumentType = "INCOME_PROOF"
	DocumentTypeCreditReport   DocumentType = "CREDIT_REPORT"
	DocumentTypeBackgroundCheck DocumentType = "BACKGROUND_CHECK"
	DocumentTypeLeaseAgreement DocumentType = "LEASE_AGREEMENT"
	DocumentTypeOther          DocumentType = "OTHER"
)

type Document struct {
	ID            string       `gorm:"type:uuid;primary_key" json:"id"`
	Name          string       `gorm:"not null" json:"name"`
	Type          DocumentType `gorm:"not null" json:"type"`
	URL           string       `gorm:"not null" json:"url"`
	ApplicationID string       `gorm:"type:uuid" json:"-"`
	UploadedByID  string       `gorm:"type:uuid" json:"-"`
	UploadedBy    User         `gorm:"foreignKey:UploadedByID" json:"uploadedBy"`
	CreatedAt     time.Time    `json:"createdAt"`
}

func (d *Document) BeforeCreate(tx *gorm.DB) error {
	if d.ID == "" {
		d.ID = uuid.New().String()
	}
	return nil
}
