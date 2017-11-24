package tpls

// Page represents an SQL LIMIT statement
type Page struct {
	limit  int64
	offset int64
}

// String returns the SQL query representation of the Page
func (p *Page) Page() (int64, int64) {
	return p.limit, p.offset
}
