package tpls

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
	"unsafe"
)

// Values is a hack to the sql.Rows struct
// since the rows struct does not expose it's lastcols values, or a way to give
// a custom scanner to the Scan method.
// See issue https://github.com/golang/go/issues/22544
func rowValues(r sql.Rows) []driver.Value {
	// some ugly hack to access lastcols field
	rs := reflect.ValueOf(&r).Elem()
	rf := rs.FieldByName("lastcols")

	// overcome panic reflect.Value.Interface: cannot return value obtained from unexported field or method
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	return rf.Interface().([]driver.Value)
}
