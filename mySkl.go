package rankdb

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	P        = 0.75
	MaxLevel = 8
)

type (
	SkipList struct {
		SkpListName string
		writeLock   sync.Mutex // 写锁
		readLock    sync.Mutex // 读锁
		Nodes       []*Node
		rand        *rand.Rand
	}

	Node struct {
		RootEle *Ele
		EleNum  int
		Level   int
		// 记录走过的上层的元素，方便插入后更改
		PreviewNodes []*Ele
	}

	// Ele 设计为双向链表
	Ele struct {
		Score CompareAble
		Value interface{}
		// 根？
		RootEl bool
		// 记录属于哪一个节点的
		Node    *Node
		NextEle *Ele
		PreEle  *Ele
	}
)

// NewRootEle 创建根Ele
func NewRootEle(node *Node) *Ele {
	rootEle := &Ele{
		RootEl: true,
		Node:   node,
	}
	rootEle.NextEle = rootEle
	rootEle.PreEle = rootEle
	return rootEle
}

// NewEle 创建一个新的Ele
func NewEle(score CompareAble, val interface{}, node *Node) *Ele {
	rootEle := &Ele{
		Score:  score,
		Value:  val,
		RootEl: false,
		Node:   node,
	}
	rootEle.NextEle = rootEle
	rootEle.PreEle = rootEle
	return rootEle
}

// NewRootNode 创建根Node
func NewRootNode(level int) *Node {
	rootNode := &Node{
		EleNum: 0,
		Level:  level,
	}
	rootNode.RootEle = NewRootEle(rootNode)
	return rootNode
}

// HasEle 是否有节点
func (n *Node) HasEle() bool {
	// 判断节点的根节点的下一个是不是自己
	if n.RootEle.NextEle.RootEl {
		return false
	}
	return true
}

// 随机层数
func (s *SkipList) randomLevel() int {
	level := 1
	for float64(s.rand.Int31()&0xFFFF) < P*0xFFFF {
		level += 1
	}
	if level < MaxLevel {
		return level
	}
	return MaxLevel
}

// NewSkl 新建skl
func NewSkl(keyName string) *SkipList {
	skl := &SkipList{
		SkpListName: keyName,
		rand:        rand.New(rand.NewSource(time.Now().Unix())),
	}
	skl.Nodes = make([]*Node, MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		skl.Nodes[i] = NewRootNode(i + 1)
	}
	return skl
}

func (s *SkipList) Insert(keyName string, score CompareAble, value interface{}) {
	if len(s.Nodes) == 0 {
		s = NewSkl(keyName)
	}
	insertLevel := s.randomLevel()

	// 暂存走过的元素,记录到的是插入key比较大小的前一个值
	prevs := make([]*Ele, MaxLevel)
	for i := MaxLevel - 1; i >= 0; i-- {
		if s.Nodes[i].HasEle() {
			// 如果此层的node有元素的话,遍历此node下的元素，并且判断当前节点是否大于插入值
			for nextEle := s.Nodes[i].RootEle.NextEle; nextEle.RootEl != true; nextEle = nextEle.NextEle {
				// 判断当前节点是否大于插入值
				if c, err := nextEle.Score.Compare(score); c >= 0 {
					if err != nil {
						return
					}
					// 判断当前层数是否小于等于随机的层数
					if i <= insertLevel-1 {
						// 填入的是当前遍历元素的上一个值，方便接下来统一插入
						prevs[i] = nextEle.PreEle
						break
					}
				} else {
					if nextEle.NextEle.RootEl == true {
						if i <= insertLevel-1 {
							// 填入的是当前遍历元素的上一个值，方便接下来统一插入
							prevs[i] = nextEle
							break
						}
					}
				}
			}
		} else {
			// 没有的话进入
			if i <= insertLevel-1 {
				prevs[i] = s.Nodes[i].RootEle
			}
		}
	}

	// 遍历并替换值
	for i := 0; i < len(prevs) && prevs[i] != nil; i++ {
		newEle := NewEle(score, value, prevs[i].Node)
		tempNextEle := prevs[i].NextEle
		prevs[i].NextEle.PreEle = newEle // 设置旧的下一个元素
		prevs[i].NextEle = newEle
		// 设置新元素
		newEle.NextEle = tempNextEle
		newEle.PreEle = prevs[i]

		prevs[i].Node.EleNum++
	}
}
func (s SkipList) Get(score CompareAble) interface{} {
	if len(s.Nodes) == 0 {
		return nil
	}
	for i := MaxLevel - 1; i >= 0; i-- {
		// 如果此层的node有元素的话,遍历此node下的元素，并且判断当前节点是否大于插入值
		for nextEle := s.Nodes[i].RootEle.NextEle; nextEle.RootEl != true; nextEle = nextEle.NextEle {
			if c, err := nextEle.Score.Compare(score); c == 0 {
				if err != nil {
					return nil
				}
				return nextEle.Value
			}
		}
	}
	return nil
}

func (s *SkipList) Print() {
	for i := 0; i < len(s.Nodes); i++ {
		var eachEleKv []string
		nextEle := s.Nodes[i].RootEle.NextEle
		for nextEle.RootEl != true {
			eachEleKv = append(eachEleKv, spew.Sdump(nextEle.Value))
			nextEle = nextEle.NextEle
		}
		fmt.Println("第 ", s.Nodes[i].Level, " 层,共有元素 ", s.Nodes[i].EleNum, " 个")
		fmt.Println(strings.Join(eachEleKv, ","))
	}
}