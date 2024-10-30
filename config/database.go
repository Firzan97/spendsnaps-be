package config

import (
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database configuration
type Database struct {
	URL  string
	Name string
}

func (db *Database) Connect() error {

	// Connect
	if err := mgm.SetDefaultConfig(
		nil, "", options.Client().ApplyURI(db.URL),
	); err != nil {
		return err
	}

	log.Printf("Connected to DB")

	return nil
}
