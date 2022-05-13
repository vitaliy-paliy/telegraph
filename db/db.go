package db

import (
	"telegraph/model"
	"telegraph/store"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	DB   *gorm.DB
	Auth *store.AuthStore
	Friendship *store.FriendshipStore
}

func (c *Client) init() {
	c.Auth = store.NewAuthStore(c.DB)
	c.Friendship = store.NewFriendshipStore(c.DB)
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
	autoMigrate(db)

	// New client.
	client = &Client{DB: db}
	client.init()

	return
}

// Auto migrate keeps schema up to date.
func autoMigrate(db *gorm.DB) {
	models := []interface{}{
		&model.User{},
		&model.Friendship{},
	}

	for _, model := range models {
		db.AutoMigrate(model)
	}
}
