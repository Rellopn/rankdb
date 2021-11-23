package rankdb

import (
	"fmt"
	"log"
	"testing"
)

func Test_f64CompareAble_Compare(t *testing.T) {
	var a, b F64CompareAble
	a = NewF64CompareAble(54.123)
	b = NewF64CompareAble(32.889)
	compareResult, err := a.CompareTo(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("except: 1, got ", compareResult)

	a = NewF64CompareAble(1995.06)
	b = NewF64CompareAble(1996.07)
	compareResult, err = a.CompareTo(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("except: -1, got ", compareResult)

	a = NewF64CompareAble(2021.06)
	b = NewF64CompareAble(2021.06)
	compareResult, err = a.CompareTo(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("except: 0, got ", compareResult)

	a = NewF64CompareAble(2021.06)
	c := NewStrCompareAble("23")
	_, err = a.CompareTo(c)
	if err != nil {
		fmt.Println("except: err, got ", err)
	}
}
