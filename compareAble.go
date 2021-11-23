package rankdb

import (
	"errors"
	"fmt"
	"reflect"
)

type CompareAble interface {
	CompareTo(CompareAble) (int, error)
}

type F64CompareAble struct {
	float64
}

func NewF64CompareAble(f64 float64) F64CompareAble {
	return F64CompareAble{f64}
}

func (f F64CompareAble) CompareTo(f64 CompareAble) (int, error) {
	f64origin, ok := f64.(F64CompareAble)
	if !ok {
		f64RealKind := reflect.TypeOf(f64).Name()
		return 0, errors.New(fmt.Sprintf("Not except type, except type F64CompareAble, but got [%s]", f64RealKind))
	}
	if f.float64 > f64origin.float64 {
		return 1, nil
	} else if f.float64 < f64origin.float64 {
		return -1, nil
	} else {
		return 0, nil
	}
}

type StrCompareAble struct {
	string
}

func NewStrCompareAble(str string) StrCompareAble {
	return StrCompareAble{str}
}

func (s StrCompareAble) CompareTo(str CompareAble) (int, error) {
	strOrigin, ok := str.(StrCompareAble)
	if !ok {
		strOriginKind := reflect.TypeOf(str).Name()
		return 0, errors.New(fmt.Sprintf("Not except type, except type StrCompareAble, but got [%s]", strOriginKind))
	}
	if s.string > strOrigin.string {
		return 1, nil
	} else if s.string > strOrigin.string {
		return -1, nil
	} else {
		return 0, nil
	}
}
