// Autogenerated by github.com/posener/orm
package allorm

import (
	"fmt"
	"strings"

	"github.com/posener/orm/example"
)

func (i *TInsert) String() string {
	return fmt.Sprintf(`INSERT INTO all (%s) VALUES (%s)`,
		strings.Join(i.cols, ", "),
		qMarks(len(i.values)),
	)
}

func (o *ORM) InsertAll(p *example.All) *TInsert {
	i := o.Insert()
	i.add("int", p.Int)
	i.add("string", p.String)
	i.add("bool", p.Bool)
	return i
}

func (i *TInsert) SetInt(value int) *TInsert {
	return i.add("int", value)
}

func (i *TInsert) SetString(value string) *TInsert {
	return i.add("string", value)
}

func (i *TInsert) SetBool(value bool) *TInsert {
	return i.add("bool", value)
}
