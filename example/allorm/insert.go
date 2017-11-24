// Autogenerated by github.com/posener/orm
package allorm

import (
	"fmt"
	"strings"
	"time"

	"github.com/posener/orm/example"
)

func (i *TInsert) String() string {
	return fmt.Sprintf(`INSERT INTO 'all' (%s) VALUES (%s)`,
		strings.Join(i.cols, ", "),
		qMarks(len(i.values)),
	)
}

// InsertAll creates an INSERT statement according to the given object
func (o *ORM) InsertAll(p *example.All) *TInsert {
	i := o.Insert()
	i.add("`int`", p.Int)
	i.add("`string`", p.String)
	i.add("`bool`", p.Bool)
	i.add("`auto`", p.Auto)
	i.add("`time`", p.Time)
	i.add("`select`", p.Select)
	return i
}

// SetInt sets value for column int in the INSERT statement
func (i *TInsert) SetInt(value int) *TInsert {
	return i.add("`int`", value)
}

// SetString sets value for column string in the INSERT statement
func (i *TInsert) SetString(value string) *TInsert {
	return i.add("`string`", value)
}

// SetBool sets value for column bool in the INSERT statement
func (i *TInsert) SetBool(value bool) *TInsert {
	return i.add("`bool`", value)
}

// SetAuto sets value for column auto in the INSERT statement
func (i *TInsert) SetAuto(value int) *TInsert {
	return i.add("`auto`", value)
}

// SetTime sets value for column time in the INSERT statement
func (i *TInsert) SetTime(value time.Time) *TInsert {
	return i.add("`time`", value)
}

// SetSelect sets value for column select in the INSERT statement
func (i *TInsert) SetSelect(value int) *TInsert {
	return i.add("`select`", value)
}
