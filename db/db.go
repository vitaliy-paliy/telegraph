package db

import (
	"telegraph/model"
	"telegraph/store"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/lib/pq"
)

type Client struct {
	DB         *gorm.DB
	Enforcer   *casbin.Enforcer
	Auth       *store.AuthStore
	Friendship *store.FriendshipStore
	Messenger  *store.MessengerStore
}

func (c *Client) init() {
	c.Auth = store.NewAuthStore(c.DB)
	c.Friendship = store.NewFriendshipStore(c.DB, c.Enforcer)
	c.Messenger = store.NewMessengerStore(c.DB)
}

func Start() (client *Client, err error) {
	// Configure gorm DB.
	dsn := "host=localhost port=5432 user=paliy password=secret dbname=telegraph sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		return
	}

	// Casbin
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return
	}

	enforcer, err := casbin.NewEnforcer("config/model.conf", adapter)
	if err != nil {
		return
	}

	// New client.
	client = &Client{DB: db, Enforcer: enforcer}
	client.init()

	autoMigrate(db)

	return
}

// Auto migrate keeps schema up to date.
func autoMigrate(db *gorm.DB) {
	models := []interface{}{
		&model.User{},
		&model.Friendship{},
		&model.Message{},
	}

	for _, model := range models {
		db.AutoMigrate(model)
	}
}
