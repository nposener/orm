// Package employeeorm was auto-generated by github.com/posener/orm; DO NOT EDIT
package employeeorm

import (
	"database/sql"

	"github.com/posener/orm"
	"github.com/posener/orm/common"
	"github.com/posener/orm/dialect"

	"github.com/posener/orm/example"
)

// table is SQL table name
const table = "employee"

// createColumnsStatements are columns definitions in different dialects
var createColumnsStatements = map[string]string{
	"mysql":   "`name` TEXT, `age` INTEGER, `salary` INTEGER",
	"sqlite3": "'name' TEXT, 'age' INTEGER, 'salary' INTEGER",
}

// API is the interface of the ORM object
type API interface {
	Close() error
	Create() *CreateBuilder
	Select() *SelectBuilder
	Insert() *InsertBuilder
	Update() *UpdateBuilder
	Delete() *DeleteBuilder
	InsertEmployee(*example.Employee) *InsertBuilder
	UpdateEmployee(*example.Employee) *UpdateBuilder

	Logger(orm.Logger)
}

// Querier is the interface for a SELECT SQL statement
type Querier interface {
	Query() ([]example.Employee, error)
}

// Counter is the interface for a SELECT SQL statement for counting purposes
type Counter interface {
	Count() ([]EmployeeCount, error)
}

// Firster is the interface for a SELECT SQL statement for getting only the
// first item. if no item matches the query, an `orm.ErrNotFound` will be returned.
type Firster interface {
	First() (*example.Employee, error)
}

// Open opens database connection
func Open(driverName, dataSourceName string) (API, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	d, err := dialect.New(driverName)
	if err != nil {
		return nil, err
	}
	return &conn{dialect: d, db: db}, nil
}

// New returns an conn object from a db instance
func New(driverName string, db orm.DB) (API, error) {
	d, err := dialect.New(driverName)
	if err != nil {
		return nil, err
	}
	return &conn{dialect: d, db: db}, nil
}

// Create returns a builder of an SQL CREATE statement
func (c *conn) Create() *CreateBuilder {
	return &CreateBuilder{
		params: common.CreateParams{
			Table:            table,
			ColumnsStatement: createColumnsStatements[c.dialect.Name()],
		},
		conn: c,
	}
}

// Select returns a builder of an SQL SELECT statement
func (c *conn) Select() *SelectBuilder {
	s := &SelectBuilder{
		params: common.SelectParams{Table: table},
		conn:   c,
	}
	s.params.Columns = &s.selector
	return s
}

// Insert returns a builder of an SQL INSERT statement
func (c *conn) Insert() *InsertBuilder {
	return &InsertBuilder{
		params: common.InsertParams{Table: table},
		conn:   c,
	}
}

// InsertEmployee returns an SQL INSERT statement builder filled with values of a given object
func (c *conn) InsertEmployee(p *example.Employee) *InsertBuilder {
	i := c.Insert()
	i.params.Assignments.Add("name", p.Name)
	i.params.Assignments.Add("age", p.Age)
	i.params.Assignments.Add("salary", p.Salary)
	return i
}

// Update returns a builder of an SQL UPDATE statement
func (c *conn) Update() *UpdateBuilder {
	return &UpdateBuilder{
		params: common.UpdateParams{Table: table},
		conn:   c,
	}
}

// UpdateEmployee returns an SQL UPDATE statement builder filled with values of a given object
func (c *conn) UpdateEmployee(p *example.Employee) *UpdateBuilder {
	u := c.Update()
	u.params.Assignments.Add("name", p.Name)
	u.params.Assignments.Add("age", p.Age)
	u.params.Assignments.Add("salary", p.Salary)
	return u
}

// Delete returns a builder of an SQL DELETE statement
func (c *conn) Delete() *DeleteBuilder {
	return &DeleteBuilder{
		params: common.DeleteParams{Table: table},
		conn:   c,
	}
}

// SetName sets value for column name in the INSERT statement
func (i *InsertBuilder) SetName(value string) *InsertBuilder {
	i.params.Assignments.Add("name", value)
	return i
}

// SetName sets value for column name in the UPDATE statement
func (u *UpdateBuilder) SetName(value string) *UpdateBuilder {
	u.params.Assignments.Add("name", value)
	return u
}

// SetAge sets value for column age in the INSERT statement
func (i *InsertBuilder) SetAge(value int) *InsertBuilder {
	i.params.Assignments.Add("age", value)
	return i
}

// SetAge sets value for column age in the UPDATE statement
func (u *UpdateBuilder) SetAge(value int) *UpdateBuilder {
	u.params.Assignments.Add("age", value)
	return u
}

// SetSalary sets value for column salary in the INSERT statement
func (i *InsertBuilder) SetSalary(value int) *InsertBuilder {
	i.params.Assignments.Add("salary", value)
	return i
}

// SetSalary sets value for column salary in the UPDATE statement
func (u *UpdateBuilder) SetSalary(value int) *UpdateBuilder {
	u.params.Assignments.Add("salary", value)
	return u
}
