package db

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

const DSN = "host=localhost port=5432 user=paliy password=secret dbname=telegraph sslmode=disable"

type Client struct {
	db *gorm.DB
}

func Start() (client *Client, err error) {
	db, err := gorm.Open(postgres.New(postgres.Config{DriverName: "postgres", DSN: DSN,}))
	if err != nil {
		return
	}
	
	client = &Client{db: db}

	return
}
