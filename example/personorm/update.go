// Autogenerated by github.com/posener/orm
package personorm

import (
	"fmt"

	"github.com/posener/orm/example"
)

func (u *TUpdate) String() string {
	return fmt.Sprintf(`UPDATE person SET %s %s`,
		u.assignmentList(),
		u.where.String(),
	)
}

func UpdatePerson(p *example.Person) *TUpdate {
	var u TUpdate
	u.add("name", p.Name)
	u.add("age", p.Age)
	return &u
}

func (u *TUpdate) SetName(value string) *TUpdate {
	return u.add("name", value)
}

func (u *TUpdate) SetAge(value int) *TUpdate {
	return u.add("age", value)
}