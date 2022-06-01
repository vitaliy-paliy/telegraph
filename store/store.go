package store

import (
	"gorm.io/gorm"

	"github.com/casbin/casbin/v2"
)

type AuthStore struct {
	db *gorm.DB
}

type FriendshipStore struct {
	db  *gorm.DB
	enf *casbin.Enforcer
}

func NewAuthStore(db *gorm.DB) *AuthStore { return &AuthStore{db: db} }

func NewFriendshipStore(db *gorm.DB, enf *casbin.Enforcer) *FriendshipStore {
	return &FriendshipStore{db: db, enf: enf}
}
