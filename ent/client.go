// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/knabben/ggql/ent/migrate"

	"github.com/knabben/ggql/ent/objecttype"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ObjectType is the client for interacting with the ObjectType builders.
	ObjectType *ObjectTypeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:     c,
		Schema:     migrate.NewSchema(c.driver),
		ObjectType: NewObjectTypeClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config:     cfg,
		ObjectType: NewObjectTypeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ObjectType.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:     cfg,
		Schema:     migrate.NewSchema(cfg.driver),
		ObjectType: NewObjectTypeClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// ObjectTypeClient is a client for the ObjectType schema.
type ObjectTypeClient struct {
	config
}

// NewObjectTypeClient returns a client for the ObjectType from the given config.
func NewObjectTypeClient(c config) *ObjectTypeClient {
	return &ObjectTypeClient{config: c}
}

// Create returns a create builder for ObjectType.
func (c *ObjectTypeClient) Create() *ObjectTypeCreate {
	return &ObjectTypeCreate{config: c.config}
}

// Update returns an update builder for ObjectType.
func (c *ObjectTypeClient) Update() *ObjectTypeUpdate {
	return &ObjectTypeUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *ObjectTypeClient) UpdateOne(ot *ObjectType) *ObjectTypeUpdateOne {
	return c.UpdateOneID(ot.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *ObjectTypeClient) UpdateOneID(id int) *ObjectTypeUpdateOne {
	return &ObjectTypeUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for ObjectType.
func (c *ObjectTypeClient) Delete() *ObjectTypeDelete {
	return &ObjectTypeDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ObjectTypeClient) DeleteOne(ot *ObjectType) *ObjectTypeDeleteOne {
	return c.DeleteOneID(ot.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ObjectTypeClient) DeleteOneID(id int) *ObjectTypeDeleteOne {
	return &ObjectTypeDeleteOne{c.Delete().Where(objecttype.ID(id))}
}

// Create returns a query builder for ObjectType.
func (c *ObjectTypeClient) Query() *ObjectTypeQuery {
	return &ObjectTypeQuery{config: c.config}
}

// Get returns a ObjectType entity by its id.
func (c *ObjectTypeClient) Get(ctx context.Context, id int) (*ObjectType, error) {
	return c.Query().Where(objecttype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ObjectTypeClient) GetX(ctx context.Context, id int) *ObjectType {
	ot, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ot
}