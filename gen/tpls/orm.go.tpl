package {{.Package}}

import (
    "database/sql"

	"github.com/posener/orm/common"
	"github.com/posener/orm/dialect"
)

const table = "{{.Type.Table}}"

var createColumnsStatements = map[string]string{
    {{ range $_, $d := .Dialects -}}
    "{{$d.Name}}": "{{$d.ColumnsStatement}}",
    {{ end -}}
}

// Open opens database connection
func Open(driverName, dataSourceName string) (API, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	dialect, err := dialect.New(driverName)
	if err != nil {
		return nil, err
	}
	return &orm{dialect: dialect, db: db}, nil
}

// New returns an orm object from a db instance
func New(driverName string, db DB) (API, error) {
	dialect, err := dialect.New(driverName)
	if err != nil {
		return nil, err
	}
    return &orm{dialect: dialect, db: db}, nil
}

// Create returns a struct for a CREATE statement
func (o *orm) Create() *Create {
	return &Create{
		internal: common.Create{
		    Table: table,
		    ColumnsStatement: createColumnsStatements[o.dialect.Name()],
        },
	    orm: o,
    }
}

// Select returns an object to create a SELECT statement
func (o *orm) Select() *Select {
	s := &Select{
		internal: common.Select{Table: table},
		orm: o,
	}
    s.internal.Columns = &s.columns
    return s
}

// Insert returns a new INSERT statement
func (o *orm) Insert() *Insert {
	return &Insert{
		internal: common.Insert{Table: table},
		orm: o,
	}
}

// Update returns a new UPDATE statement
func (o *orm) Update() *Update {
	return &Update{
		internal: common.Update{Table: table},
		orm: o,
    }
}

// Delete returns an object for a DELETE statement
func (o *orm) Delete() *Delete {
	return &Delete{
		internal: common.Delete{Table: table},
		orm: o,
    }
}

