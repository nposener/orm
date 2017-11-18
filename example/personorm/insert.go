// Autogenerated by github.com/posener/orm
package personorm

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/posener/orm/tools"

	"github.com/posener/orm/example"
)

func NewInsert() Insert {
	return Insert{}
}

type Insert struct {
	cols   []string
	values []interface{}
}

func (i Insert) Name(value string) Insert {
	return i.add("name", value)
}

func (i Insert) Age(value int) Insert {
	return i.add("age", value)
}

func (i Insert) Person(p *example.Person) Insert {
	var j = i
	j = j.add("name", p.Name)
	j = j.add("age", p.Age)
	return j
}

// Create creates a table for example.Person
func (i Insert) Exec(db *sql.DB) error {
	if len(i.cols) == 0 || len(i.values) == 0 {
		return fmt.Errorf("nothing to insert")
	}
	stmt := fmt.Sprintf(`INSERT INTO person (%s) VALUES (%s)`,
		strings.Join(i.cols, ", "),
		tools.QMarks(len(i.values)),
	)

	log.Printf("Insert: '%v' (%v)", stmt, i.values)
	_, err := db.Exec(stmt, i.values...)
	return err
}

func (i Insert) add(name string, value interface{}) Insert {
	i.cols = append(i.cols, name)
	i.values = append(i.values, value)
	return i
}
