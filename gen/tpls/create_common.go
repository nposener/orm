package tpls

import (
	"database/sql"
	"fmt"
	"log"
)

// TCreate is a struct that holds data for the CREATE statement
type TCreate struct {
	fmt.Stringer
	sql.Result
}

// Create returns a struct for a CREATE statement
func Create() *TCreate {
	return &TCreate{}
}

// Exec creates a table for the given struct
func (c *TCreate) Exec(db SQLExecer) error {
	stmt := c.String()
	log.Printf("Create: '%v'", stmt)
	result, err := db.Exec(stmt)
	c.Result = result
	return err
}
