package rankdb

import (
	"fmt"
	"testing"
)

func Test_canCompareRanElements_RetainAll(t *testing.T) {
	a1 := &canCompareRanElement{Rank: 1, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	a2 := &canCompareRanElement{Rank: 2, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	a3 := &canCompareRanElement{Rank: 3, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	a4 := &canCompareRanElement{Rank: 4, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	a5 := &canCompareRanElement{Rank: 5, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	a6 := &canCompareRanElement{Rank: 6, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	a5.Next = a6

	b1 := &canCompareRanElement{Rank: 1, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	b2 := &canCompareRanElement{Rank: 2, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	b5 := &canCompareRanElement{Rank: 5, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	b6 := &canCompareRanElement{Rank: 6, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}
	b7 := &canCompareRanElement{Rank: 7, Pointer: &RaceRankResult{Name: "mark", Age: 23, ShoeBrand: "Nk", Speed: 1515.12357}}

	b1.Next = b2
	b2.Next = b5
	b5.Next = b6
	b6.Next = b7

	fmt.Println("test 1,2")
	// test
	// 1. |________|
	//        |____|
	// ===============
	// 2.     |____|
	//    |________|
	a := canCompareRanElements{a2, a3, a4, a5, a6}
	b := canCompareRanElements{b2, b5, b6}
	getRetainAllc := a.RetainAll(b)
	for _, g := range getRetainAllc {
		fmt.Println(g.Rank)
	}
	fmt.Println("---------------------------")
	getRetainAlld := b.RetainAll(a)
	for _, gb := range getRetainAlld {
		fmt.Println(gb.Rank)
	}
	fmt.Println("test 3,4")
	// test
	// 3. |________|
	//        |____|
	// ===============
	// 4.     |____|
	//    |________|
	a = canCompareRanElements{a1, a2, a3, a4, a5, a6}
	b = canCompareRanElements{b1, b2, b5}
	getRetainAllc = a.RetainAll(b)
	for _, g := range getRetainAllc {
		fmt.Println(g.Rank)
	}
	fmt.Println("---------------------------")
	getRetainAlld = b.RetainAll(a)
	for _, gb := range getRetainAlld {
		fmt.Println(gb.Rank)
	}
	fmt.Println("test 5,6")
	// test
	// 3.    |________|
	//    |____|
	// ===============
	// 4.    |________|
	//    |____|
	a = canCompareRanElements{a2, a3, a4, a5, a6}
	b = canCompareRanElements{b1, b2, b5}
	getRetainAllc = a.RetainAll(b)
	for _, g := range getRetainAllc {
		fmt.Println(g.Rank)
	}
	fmt.Println("---------------------------")
	getRetainAlld = b.RetainAll(a)
	for _, gb := range getRetainAlld {
		fmt.Println(gb.Rank)
	}
}
