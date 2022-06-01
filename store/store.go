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

type MessengerStore struct {
	db *gorm.DB
}

func NewAuthStore(db *gorm.DB) *AuthStore { return &AuthStore{db: db} }

func NewFriendshipStore(db *gorm.DB, enf *casbin.Enforcer) *FriendshipStore {
	return &FriendshipStore{db: db, enf: enf}
}

func NewMessengerStore(db *gorm.DB) *MessengerStore { return &MessengerStore{db: db} }
