package models

import (
	"gorm.io/gorm"
	"time"
)

type Otp struct {
	ID        int64          `gorm:"type:primaryKey" json:"id"`
	Email     string         `gorm:"unique" json:"email"`
	Otp       string         `json:"otp"`
	ExpiresAt time.Time      `json:"expires_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (otp *Otp) BeforeCreate(tx *gorm.DB) (err error) {
	otp.ID = time.Now().Unix()
	otp.ExpiresAt = time.Now().Add(time.Minute * 3)
	return nil
}
