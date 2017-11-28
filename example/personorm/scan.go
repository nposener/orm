// Package personorm was auto-generated by github.com/posener/orm; DO NOT EDIT
package personorm

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"unsafe"
)

const errMsg = "converting %s: column %d with value %v (type %T) to %s"

// scanArgs are list of fields to be given to the sql Scan command
func (c *columns) scan(dialect string, rows *sql.Rows) (*PersonCount, error) {
	switch dialect {
	case "mysql":
		return c.scanmysql(rows)

	case "sqlite3":
		return c.scansqlite3(rows)
	default:
		return nil, fmt.Errorf("unsupported dialect %s", dialect)
	}
}

func (c *columns) scanmysql(rows *sql.Rows) (*PersonCount, error) {
	var (
		vals = values(*rows)
		row  PersonCount
		all  = c.selectAll()
		i    = 0
	)

	if all || c.SelectName {
		if vals[i] != nil {
			switch val := vals[i].(type) {
			case []byte:
				tmp := string(val)
				row.Name = tmp
			default:
				return nil, fmt.Errorf(errMsg, "Name", i, vals[i], vals[i], "[]byte, []byte")
			}
		}
		i++
	}

	if all || c.SelectAge {
		if vals[i] != nil {
			switch val := vals[i].(type) {
			case []byte:
				tmp := int(parseInt(val))
				row.Age = tmp
			case int64:
				tmp := int(val)
				row.Age = tmp
			default:
				return nil, fmt.Errorf(errMsg, "Age", i, vals[i], vals[i], "[]byte, int64")
			}
		}
		i++
	}

	if c.count {
		switch val := vals[i].(type) {
		case int64:
			row.Count = val
		case []byte:
			row.Count = parseInt(val)
		default:
			return nil, fmt.Errorf(errMsg, "COUNT(*)", i, vals[i], vals[i], "int64, []byte")
		}
	}

	return &row, nil
}

func (c *columns) scansqlite3(rows *sql.Rows) (*PersonCount, error) {
	var (
		vals = values(*rows)
		row  PersonCount
		all  = c.selectAll()
		i    = 0
	)

	if all || c.SelectName {
		if vals[i] != nil {
			val, ok := vals[i].([]byte)
			if !ok {
				return nil, fmt.Errorf(errMsg, "Name", i, vals[i], vals[i], "string")
			}
			tmp := string(val)
			row.Name = tmp
		}
		i++
	}

	if all || c.SelectAge {
		if vals[i] != nil {
			val, ok := vals[i].(int64)
			if !ok {
				return nil, fmt.Errorf(errMsg, "Age", i, vals[i], vals[i], "int")
			}
			tmp := int(val)
			row.Age = tmp
		}
		i++
	}

	if c.count {
		switch val := vals[i].(type) {
		case int64:
			row.Count = val
		case []byte:
			row.Count = parseInt(val)
		default:
			return nil, fmt.Errorf(errMsg, "COUNT(*)", i, vals[i], vals[i], "int64, []byte")
		}
	}

	return &row, nil
}

// Values is a hack to the sql.Rows struct
// since the rows struct does not expose it's lastcols values, or a way to give
// a custom scanner to the Scan method.
// See issue https://github.com/golang/go/issues/22544
func values(r sql.Rows) []driver.Value {
	// some ugly hack to access lastcols field
	rs := reflect.ValueOf(&r).Elem()
	rf := rs.FieldByName("lastcols")

	// overcome panic reflect.Value.Interface: cannot return value obtained from unexported field or method
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	return rf.Interface().([]driver.Value)
}
