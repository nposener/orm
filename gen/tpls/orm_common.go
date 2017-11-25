package tpls

import (
	"context"
	"database/sql"

	"github.com/posener/orm/common"
	"github.com/posener/orm/dialect"
)

// DB is an interface of functions of sql.DB which are used by orm struct.
type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	Close() error
}

// orm represents an orm of a given struct.
// All functions available to interact with an SQL table that is related
// to this struct, are done by an instance of this object.
// To get an instance of orm use Open or New functions.
type orm struct {
	dialect dialect.Dialect
	db      DB
	logger  Logger
}

func (o *orm) Close() error {
	return o.db.Close()
}

// Logger sets a logger to the orm package
func (o *orm) Logger(logger Logger) {
	o.logger = logger
}

// Create is a struct that holds data for the CREATE statement
type Create struct {
	internal common.Create
	orm      *orm
}

// IfNotExists sets IF NOT EXISTS for the CREATE SQL statement
func (c *Create) IfNotExists() *Create {
	c.internal.IfNotExists = true
	return c
}

// Insert is a struct to hold information for an INSERT statement
type Insert struct {
	internal common.Insert
	orm      *orm
}

// Update is a struct to hold information for an INSERT statement
type Update struct {
	internal common.Update
	orm      *orm
}

func (u *Update) Where(where common.Where) *Update {
	u.internal.Where = where
	return u
}

// Delete is the struct that holds the SELECT data
type Delete struct {
	internal common.Delete
	orm      *orm
}

// Where applies where conditions on the query
func (d *Delete) Where(w common.Where) *Delete {
	d.internal.Where = w
	return d
}
