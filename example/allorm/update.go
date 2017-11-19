// Autogenerated by github.com/posener/orm
package allorm

import (
	"fmt"

	"github.com/posener/orm/example"
)

func (u *TUpdate) String() string {
	return fmt.Sprintf(`UPDATE all SET %s %s`,
		u.assignmentList(),
		u.where.String(),
	)
}

func UpdateAll(p *example.All) *TUpdate {
	var u TUpdate
	u.add("int", p.Int)
	u.add("text", p.Text)
	u.add("bool", p.Bool)
	return &u
}

func (u *TUpdate) SetInt(value int) *TUpdate {
	return u.add("int", value)
}

func (u *TUpdate) SetText(value string) *TUpdate {
	return u.add("text", value)
}

func (u *TUpdate) SetBool(value bool) *TUpdate {
	return u.add("bool", value)
}