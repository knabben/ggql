package database

import (
	"github.com/knabben/ggql/ent"
)

type Database struct {
	dbURI string
	client *ent.Client
}

// NewDatabase create a new schema and connect
func NewDatabase(uri string) (*Database, error) {
	db := &Database{dbURI: uri}
	client, err := db.Connect()

	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}
	db.client = client
	return db, nil
}

func (d *Database) Connect() (*ent.Client, error) {
	c, err := ent.Open("sqlite3", d.dbURI)
	if err != nil {
		return nil, err
	}
	return c, nil
}
