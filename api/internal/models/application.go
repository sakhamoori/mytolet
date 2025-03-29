package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplicationStatus string

const (
	ApplicationStatusSubmitted ApplicationStatus = "SUBMITTED"
	ApplicationStatusScreening ApplicationStatus = "SCREENING"
	ApplicationStatusApproved  ApplicationStatus = "APPROVED"
	ApplicationStatusRejected  ApplicationStatus = "REJECTED"
	ApplicationStatusWithdrawn ApplicationStatus = "WITHDRAWN"
)

type Application struct {
	ID          string            `gorm:"type:uuid;primary_key" json:"id"`
	PropertyID  string            `gorm:"type:uuid" json:"-"`
	Property    Property          `gorm:"foreignKey:PropertyID" json:"property"`
	ApplicantID string            `gorm:"type:uuid" json:"-"`
	Applicant   User              `gorm:"foreignKey:ApplicantID" json:"applicant"`
	Status      ApplicationStatus `gorm:"not null" json:"status"`
	CreditScore int               `json:"creditScore"`
	Income      float64           `json:"income"`
	Documents   []Document        `gorm:"foreignKey:ApplicationID" json:"documents"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

func (a *Application) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	if a.Status == "" {
		a.Status = ApplicationStatusSubmitted
	}
	return nil
}