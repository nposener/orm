// Autogenerated by github.com/posener/orm; DO NOT EDIT
package personsqlite3

import "github.com/posener/orm/dialect/sqlite3"

// String returns the SQL DELETE statement
func (s *TDelete) String() string {
	return sqlite3.Delete(s.orm, s.where)
}
