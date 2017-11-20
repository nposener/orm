// Autogenerated by github.com/posener/orm
package allorm

import "github.com/posener/orm/example"

type API interface {
	Create() *TCreate
	Query() *Select
	Insert() *TInsert
	Update() *TUpdate
	Delete() *TDelete
	InsertAll(*example.All) *TInsert
	UpdateAll(*example.All) *TUpdate
}