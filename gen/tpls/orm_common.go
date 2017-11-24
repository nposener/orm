package tpls

import (
	"database/sql"

	"github.com/posener/orm/dialect/api"
)

type ORM struct {
	dialect api.Dialect
	db      *sql.DB
	logger  Logger
}

func (o *ORM) Close() error {
	return o.db.Close()
}

// Create returns a struct for a CREATE statement
func (o *ORM) Create() *TCreate {
	return &TCreate{orm: o}
}

// Insert returns a new INSERT statement
func (o *ORM) Insert() *TInsert {
	return &TInsert{orm: o}
}

// Insert returns a new INSERT statement
func (o *ORM) Update() *TUpdate {
	return &TUpdate{TInsert: TInsert{orm: o}}
}

// Delete returns an object for a DELETE statement
func (o *ORM) Delete() *TDelete {
	return &TDelete{orm: o}
}

// Logger sets a logger to the ORM package
func (o *ORM) Logger(logger Logger) {
	o.logger = logger
}
