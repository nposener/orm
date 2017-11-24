package tpls

import (
	"database/sql"
	"fmt"

	"github.com/posener/orm/dialect/api"
)

// TInsert is a struct to hold information for an INSERT statement
type TInsert struct {
	Execer
	Argser
	fmt.Stringer
	orm    *ORM
	assign []api.Assignment
}

// Exec inserts the data to the given database
func (i *TInsert) Exec() (sql.Result, error) {
	if len(i.assign) == 0 {
		return nil, fmt.Errorf("nothing to insert")
	}

	stmt := i.String()
	args := i.Args()
	i.orm.log("Insert: '%v' %v", stmt, args)
	return i.orm.db.Exec(stmt, args...)
}

// Args returns a list of arguments for the INSERT statement
func (i *TInsert) Args() []interface{} {
	args := make([]interface{}, len(i.assign))
	for j := range i.assign {
		args[j] = i.assign[j].Value
	}
	return args
}

func (i *TInsert) add(name string, value interface{}) *TInsert {
	i.assign = append(i.assign, api.Assignment{Column: name, Value: value})
	return i
}
