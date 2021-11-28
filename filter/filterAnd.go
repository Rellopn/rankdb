package filter

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

func (and filterAnd) WriteTo(Writer) error {
	return nil
}

func (and filterAnd) Build(Writer) error {
	return nil
}
