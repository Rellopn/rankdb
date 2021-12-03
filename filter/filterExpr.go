package filter

import "fmt"

type expr struct {
	filter Filter
}

var _ Filter = expr{}

// Expr generate customerize SQL
func Expr(filter Filter) Filter {
	return expr{filter: filter}
}

func (expr expr) OpWriteTo(op string, w Writer) error {
	return expr.WriteTo(w)
}

func (expr expr) WriteTo(w Writer) error {
	if _, err := fmt.Fprint(w, "("); err != nil {
		return err
	}
	return nil
}

func (expr expr) And(conds ...Filter) Filter {
	return And(expr, And(conds...))
}

func (expr expr) Or(conds ...Filter) Filter {
	return Or(expr, Or(conds...))
}

func (expr expr) IsValid() bool {
	return expr.filter.IsValid()
}
