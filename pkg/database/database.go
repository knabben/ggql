package database

import (
	"github.com/knabben/ggql/ent"
)

type Database struct {
	dbURI string
}

var uri = "file:ent1?cache=shared&_fk=1"

func NewDatabase(uri string) *Database {
	return &Database{dbURI: uri}
}

func (d *Database) Connect() (*ent.Client, error) {
	c, err := ent.Open("sqlite3", d.dbURI)
	if err != nil {
		return nil, err
	}
	return c, nil
}
