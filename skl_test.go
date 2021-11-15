package rankdb

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	l := NewSkipList()
	for i := 0; i < 18; i++ {
		l.Put([]byte(strconv.Itoa(i)), strconv.Itoa(i))
	}
	fmt.Println(l)
}
