package store

import (
	"gorm.io/gorm"
)

type AuthStore struct {
	db *gorm.DB
}

func NewAuthStore(db *gorm.DB) *AuthStore { return &AuthStore{db: db} }
