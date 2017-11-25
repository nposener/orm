package {{.Package}}

import (
	"context"
	"database/sql"
	"fmt"
)

// Exec creates a table for the given struct
func (c *Create) Exec(ctx context.Context) (sql.Result, error) {
	stmt, args := c.orm.dialect.Create(&c.internal)
	c.orm.log("Create: '%v' %v", stmt, args)
	return c.orm.db.ExecContext(ctx, stmt, args...)
}

// query is used by the Select.Query and Select.Limit functions
func (s *Select) query(ctx context.Context) (*sql.Rows, error) {
	stmt, args := s.orm.dialect.Select(&s.internal)
	s.orm.log("Query: '%v' %v", stmt, args)
	return s.orm.db.QueryContext(ctx, stmt, args...)
}

// Exec inserts the data to the given database
func (i *Insert) Exec(ctx context.Context) (sql.Result, error) {
	if len(i.internal.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to insert")
	}
	stmt, args := i.orm.dialect.Insert(&i.internal)
	i.orm.log("Insert: '%v' %v", stmt, args)
	return i.orm.db.ExecContext(ctx, stmt, args...)
}

// Exec inserts the data to the given database
func (u *Update) Exec(ctx context.Context) (sql.Result, error) {
	if len(u.internal.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}
	stmt, args := u.orm.dialect.Update(&u.internal)
	u.orm.log("Update: '%v' %v", stmt, args)
	return u.orm.db.ExecContext(ctx, stmt, args...)
}

// Exec runs the delete statement on a given database.
func (d *Delete) Exec(ctx context.Context) (sql.Result, error) {
	stmt, args := d.orm.dialect.Delete(&d.internal)
	d.orm.log("Delete: '%v' %v", stmt, args)
	return d.orm.db.ExecContext(ctx, stmt, args...)
}
