// Autogenerated by github.com/posener/orm
package allorm

import (
	"database/sql"
	"log"

	"github.com/posener/orm/example"
)

// String returns the SQL query string
func (q *TQuery) String() string {
	return "SELECT " + q.sel.String() + " FROM all " + q.where.String()
}

// Exec runs the Query on a given database.
func (q *TQuery) Exec(db *sql.DB) ([]example.All, error) {
	// create select statement
	stmt := q.String()
	args := q.where.Args()
	log.Printf("Query: '%v' %v", stmt, args)
	rows, err := db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract rows to structures
	var all []example.All
	for rows.Next() {
		var i example.All
		if err := rows.Scan(q.sel.scanArgs(&i)...); err != nil {
			return nil, err
		}
		all = append(all, i)
	}
	return all, rows.Err()
}