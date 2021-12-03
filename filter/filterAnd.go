package filter

import "fmt"

type filterAnd []Filter

var _ Filter = filterAnd{}

// And generates AND conditions
func And(filters ...Filter) Filter {
	var result = make(filterAnd, 0, len(filters))
	for _, filter := range filters {
		if filter == nil || !filter.IsValid() {
			continue
		}
		result = append(result, filter)
	}
	return result
}

func (and filterAnd) IsValid() bool {
	return len(and) > 0
}

func (and filterAnd) And(filters ...Filter) Filter {
	return And(and, And(filters...))
}

func (and filterAnd) Or(filters ...Filter) Filter {
	return Or(and, Or(filters...))
}

func (and filterAnd) WriteTo(w Writer) error {
	for i, cond := range and {
		_, isOr := cond.(filterOr)
		_, isExpr := cond.(expr)
		wrap := isOr || isExpr
		if wrap {
			fmt.Fprint(w, "(")
		}

		err := cond.WriteTo(w)
		if err != nil {
			return err
		}

		if wrap {
			fmt.Fprint(w, ")")
		}

		if i != len(and)-1 {
			fmt.Fprint(w, " AND ")
		}
	}

	return nil
}

func (and filterAnd) Build(Writer) error {
	return nil
}
