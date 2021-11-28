package filter

type filterOr []Filter

var _ Filter = filterOr{}

// Or sets OR conditions
func Or(filters ...Filter) Filter {
	var result = make(filterOr, 0, len(filters))
	for _, Filter := range filters {
		if Filter == nil || !Filter.IsValid() {
			continue
		}
		result = append(result, Filter)
	}
	return result
}

// WriteTo implments Filter
func (o filterOr) WriteTo(w Writer) error {
	return nil
}

func (o filterOr) And(filters ...Filter) Filter {
	return And(o, And(filters...))
}

func (o filterOr) Or(filters ...Filter) Filter {
	return Or(o, Or(filters...))
}

func (o filterOr) IsValid() bool {
	return len(o) > 0
}
