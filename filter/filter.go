package filter

import "io"

type Writer interface {
	io.Writer
	Append(...interface{})
}

type Filter interface {
	WriteTo(Writer) error
	And(...Filter) Filter
	Or(...Filter) Filter
	IsValid() bool
}
