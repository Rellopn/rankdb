package rankdb

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestNewSkl(t *testing.T) {
	skl := NewSkl()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		n := r.Intn(100)
		skl.Insert(NewF64CompareAble(float64(n)), "number "+strconv.Itoa(n))
	}
	skl.Print()
	fmt.Println("-------------------")
	fmt.Println(skl.Get(NewF64CompareAble(float64(1))))
	fmt.Println(skl.Get(NewF64CompareAble(float64(4))))
	fmt.Println(skl.Get(NewF64CompareAble(float64(2))))
	fmt.Println(skl.Get(NewF64CompareAble(float64(10))))
}
