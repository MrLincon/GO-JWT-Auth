package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	ID        int64          `gorm:"type:primaryKey" json:"id"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (auth *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	auth.ID = time.Now().Unix()
	if len(auth.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		auth.Password = string(hashedPassword)
	}
	return nil
}
