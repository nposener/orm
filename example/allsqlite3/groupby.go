// Autogenerated by github.com/posener/orm; DO NOT EDIT
package allsqlite3

import (
	"github.com/posener/orm/dialect/api"
)

type groupBy []api.Group

// Add adds a column to the grouping
func (g *groupBy) add(column string) {
	*g = append(*g, api.Group{Column: column})
}
