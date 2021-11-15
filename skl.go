package rankdb

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	// the max level of the skl indexes, can be adjusted according to the actual situation.
	maxLevel    int     = 18
	probability float64 = 1 / math.E
)

type (
	Node struct {
		next []*Element
	}

	Element struct {
		Node
		key   []byte
		value interface{}
	}

	// SkipList define the skip list.
	SkipList struct {
		Node
		maxLevel       int
		Len            int
		randSource     rand.Source
		probability    float64
		probTable      []float64
		prevNodesCache []*Node
	}
)

// NewSkipList create a new skip list.
func NewSkipList() *SkipList {
	return &SkipList{
		Node:           Node{next: make([]*Element, maxLevel)},
		prevNodesCache: make([]*Node, maxLevel),
		maxLevel:       maxLevel,
		randSource:     rand.New(rand.NewSource(time.Now().UnixNano())),
		probability:    probability,
		probTable:      probabilityTable(probability, maxLevel),
	}
}

func probabilityTable(probability float64, maxLevel int) (table []float64) {
	for i := 1; i <= maxLevel; i++ {
		prob := math.Pow(probability, float64(i-1))
		table = append(table, prob)
	}
	return table
}

// Key the key of the Element.
func (e *Element) Key() []byte {
	return e.key
}

// Value the value of the Element.
func (e *Element) Value() interface{} {
	return e.value
}

// SetValue set the element value.
func (e *Element) SetValue(val interface{}) {
	e.value = val
}

// Next the first-level index of the skip list is the original data, which is arranged in an orderly manner.
// A linked list of all data in series can be obtained according to the Next method.
func (e *Element) Next() *Element {
	return e.next[0]
}

// Front first element.
// Get the head element of skl, and get all data by traversing backward.
//	e := list.Front()
//	for p := e; p != nil; p = p.Next() {
//		//do something with Element p
//	}
func (t *SkipList) Front() *Element {
	return t.next[0]
}

// Put an element into skip list, replace the value if key already exists.
func (t *SkipList) Put(key []byte, value interface{}) *Element {
	var element *Element
	prev := t.backNodes(key)

	if element = prev[0].next[0]; element != nil && bytes.Compare(element.key, key) <= 0 {
		element.value = value
		return element
	}

	element = &Element{
		Node: Node{
			next: make([]*Element, t.randomLevel()),
		},
		key:   key,
		value: value,
	}

	for i := range element.next {
		element.next[i] = prev[i].next[i]
		prev[i].next[i] = element
	}

	t.Len++
	return element
}

// generate random index level.
func (t *SkipList) randomLevel() (level int) {
	r := float64(t.randSource.Int63()) / (1 << 63)
	//fmt.Println(r)
	level = 1
	for level < t.maxLevel && r < t.probTable[level] {
		level++
	}
	//if level >2{
	fmt.Println(level)
	//}
	return
}

// find the previous node at the key.
func (t *SkipList) backNodes(key []byte) []*Node {
	var prev = &t.Node
	var next *Element

	prevs := t.prevNodesCache

	for i := t.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]

		for next != nil && bytes.Compare(key, next.key) > 0 {
			prev = &next.Node
			next = next.next[i]
		}

		prevs[i] = prev
	}

	return prevs
}
