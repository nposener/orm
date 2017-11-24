package {{.Package}}

import "database/sql"

// Open opens database connection
func Open(dataSourceName string) (*ORM, error) {
	db, err := sql.Open("{{.Dialect.Name}}", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &ORM{db: db}, nil
}

// Select returns an object to create a SELECT statement
func (o *ORM) Select() *TSelect {
	return &TSelect{orm: o}
}

func (o *ORM) Table() string {
    return "{{.Type.Table}}"
}
