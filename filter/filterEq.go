package filter

// Eq defines equals conditions
type Eq map[string]interface{}

var _ Filter = Eq{}

// WriteTo writes SQL to Writer
func (eq Eq) WriteTo(w Writer) error {
	return nil
}

// And implements And with other conditions
func (eq Eq) And(conds ...Filter) Filter {
	return And(eq, And(conds...))
}

// Or implements Or with other conditions
func (eq Eq) Or(conds ...Filter) Filter {
	return Or(eq, Or(conds...))
}

// IsValid tests if this Eq is valid
func (eq Eq) IsValid() bool {
	return len(eq) > 0
}
