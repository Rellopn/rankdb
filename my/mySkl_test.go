package my

import (
	"fmt"
	"testing"
)

func TestRandomLevel(t *testing.T) {
	skl := NewSkl()
	//for i:=0;i<100;i++{
	//	fmt.Println(skl.randomLevel())
	//}
	skl.Insert(1, "一")
	fmt.Println(1)
	skl.Insert(4, "四")
	fmt.Println(4)
	skl.Insert(7, "七")
	fmt.Println(7)
	skl.Insert(3, "三")
	fmt.Println(3)
	skl.Insert(10, "十")
	fmt.Println(10)
	skl.Insert(6, "六")
	fmt.Println(6)
	skl.Insert(2, "二")
	fmt.Println(2)
	skl.Print()
}
