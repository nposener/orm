package tpls

import (
	"database/sql"
	"fmt"
	"log"
)

// TCreate is a struct that holds data for the CREATE statement
type TCreate struct {
	db DB
	fmt.Stringer
}

// Exec creates a table for the given struct
func (c *TCreate) Exec() (sql.Result, error) {
	stmt := c.String()
	log.Printf("Create: '%v'", stmt)
	return c.db.Exec(stmt)
}
