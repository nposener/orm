package tpls

import (
	"github.com/posener/orm/dialect/api"
)

type OrderDir string

const (
	Asc  OrderDir = "ASC"
	Desc OrderDir = "DESC"
)

type orderBy []api.Order

// Add adds a column to the grouping
func (g *orderBy) add(column string, dir OrderDir) {
	*g = append(*g, api.Order{Column: column, Dir: string(dir)})
}
