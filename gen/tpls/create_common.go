package tpls

import (
	"database/sql"
	"fmt"
	"log"
)

type TCreate struct {
	fmt.Stringer
}

func Create() TCreate {
	return TCreate{}
}

// Create creates a table for {{.Type.Name}}
func (c TCreate) Exec(db *sql.DB) error {
	stmt := c.String()
	log.Printf("Create: '%v'", stmt)
	_, err := db.Exec(stmt)
	return err
}