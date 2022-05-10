package store_test

import (
	"fmt"
	"testing"

	"telegraph/db"
	// "github.com/stretchr/testify/assert"
)

func Test_AuthStore(t *testing.T) {
	dsn := "host=localhost port=5432 user=paliy password=secret dbname=telegraph sslmode=disable"
	client, _ := db.Start(dsn)
	fmt.Println(testing)
}
