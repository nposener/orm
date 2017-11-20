package tpls

import (
	"database/sql"
	"fmt"
	"log"
)

// Select is the struct that holds the SELECT data
type TDelete struct {
	db DB
	fmt.Stringer
	where *Where
}

// Where applies where conditions on the query
func (d *TDelete) Where(w *Where) *TDelete {
	d.where = w
	return d
}

// Exec runs the delete statement on a given database.
func (d *TDelete) Exec() (sql.Result, error) {
	// create select statement
	stmt := d.String()
	args := d.where.Args()
	log.Printf("Delete: '%v' %v", stmt, args)
	return d.db.Exec(stmt, args...)
}
