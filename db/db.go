package db

import (
	"telegraph/model"
	"telegraph/store"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=localhost port=5432 user=paliy password=secret dbname=telegraph sslmode=disable"

type Client struct {
	db   *gorm.DB
	Auth *store.AuthStore
}

func (c *Client) init() {
	c.Auth = store.NewAuthStore(c.db)
}

func Start() (client *Client, err error) {
	// Configure gorm DB.
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        DSN,
	}))
	if err != nil {
		return
	}
	autoMigrate(db)

	// New client.
	client = &Client{db: db}
	client.init()

	return
}

// Auto migrate keeps schema up to date.
func autoMigrate(db *gorm.DB) {
	models := []interface{}{
		&model.User{},
	}

	for _, model := range models {
		db.AutoMigrate(model)
	}
}
