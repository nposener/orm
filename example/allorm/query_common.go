// Autogenerated by github.com/posener/orm
package allorm

// Query returns a struct to perform query operations
func Query() TQuery {
	return TQuery{}
}

type TQuery struct {
	sel   TSelect
	where Where
}

func (q TQuery) Select(s TSelect) TQuery {
	q.sel = s
	return q
}

func (q TQuery) Where(w Where) TQuery {
	q.where = w
	return q
}