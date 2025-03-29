package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID         string     `gorm:"type:uuid;primary_key" json:"id"`
	SenderID   string     `gorm:"type:uuid" json:"-"`
	Sender     User       `gorm:"foreignKey:SenderID" json:"sender"`
	ReceiverID string     `gorm:"type:uuid" json:"-"`
	Receiver   User       `gorm:"foreignKey:ReceiverID" json:"receiver"`
	Content    string     `gorm:"not null" json:"content"`
	ReadAt     *time.Time `json:"readAt"`
	CreatedAt  time.Time  `json:"createdAt"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return nil
}