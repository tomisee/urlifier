package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"log"

	"fmt"

	badger "github.com/dgraph-io/badger/v3"
)

func main() {

	// Echo instance for web-app
	e := echo.New()
	e.GET("/", func(c echo.Context) error { // Route for root, using GET method
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323")) // Start web-app on port 1323

	db, err := badger.Open(badger.DefaultOptions("data")) // Open badger database in data folder
	if err != nil {
		log.Fatal(err) // If error, log it and exit
	}
	defer db.Close() // Close database when done

	// Set key-value pair as "mykey"-"myvalue"
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("mykey"), []byte("myvalue"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	// Read the key-value pair from the database to verify it was added/updated
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("mykey"))
		if err != nil {
			return err
		}
		value, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		fmt.Printf("key=%s, value=%s\n", "mykey", value)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
