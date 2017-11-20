// Autogenerated by github.com/posener/orm
package allorm

import (
	"database/sql"
	"fmt"
)

// Execer is the interface for SQL update operations
type Execer interface {
	fmt.Stringer
	Exec() (sql.Result, error)
}