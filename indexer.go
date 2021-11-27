package rankdb

import (
	"sort"
	"strings"
)

const (
	IndexTypeHash = iota
	IndexTypeRedBlackTree
)

type IndexerHash struct {
	IdxHash map[string]*OneIndexer
}

func NewIndexerHash() *IndexerHash {
	return &IndexerHash{IdxHash: make(map[string]*OneIndexer)}
}

// Indexer 实现了从sort set value中提取字段作为索引
type Indexer struct {
	// 字段的名称
	ColumnValue interface{}
	// 字段所对应的 结构体的指针
	Pointers []*PointersIndex
}

// PointersIndex 索引的Pointers的再sortSet中的排名
type PointersIndex struct {
	Pointer interface{}
	Rank    int64 // 再sortSet中的排序
}

func (i *Indexer) appendPointers(pointer *PointersIndex) {
	i.Pointers = append(i.Pointers, pointer)
}

// OneIndexer 一条索引的全部记录
type OneIndexer struct {
	IndexType int
	// 记录重复值在数组中的位置，用拉链法
	storeColumnValues map[interface{}]int
	FullIndexer       []*Indexer
}

func NewOneIndexer() *OneIndexer {
	return &OneIndexer{FullIndexer: make([]*Indexer, 0), storeColumnValues: make(map[interface{}]int)}
}

func (o *OneIndexer) StoreColumnValueExit(check interface{}) int {
	return o.storeColumnValues[check]
}

func (o *OneIndexer) ZipperAdd(columnValues, pointer interface{}, rank int64) {
	if zipperIndex := o.StoreColumnValueExit(columnValues); zipperIndex != 0 {
		// 存在的话添加数组元素
		o.FullIndexer[zipperIndex-1].appendPointers(&PointersIndex{Pointer: pointer, Rank: rank})
	} else {
		o.storeColumnValues[columnValues] = 1
		o.FullIndexer = append(o.FullIndexer,
			&Indexer{ColumnValue: columnValues, Pointers: []*PointersIndex{
				{Pointer: pointer, Rank: rank},
			}})
	}
}

func (o *OneIndexer) getFullIndexerLen() int {
	return len(o.FullIndexer)
}

// GetMemberLikeByStr 模糊搜索
func (o *OneIndexer) GetMemberLikeByStr(likeStr string) canCompareRanElements {
	var searchResult []*canCompareRanElement
	// 排名在数组中位置
	var rankInResIndex = make(map[int64]*canCompareRanElement)
	for i := 0; i < len(o.FullIndexer); i++ {
		onIndexTotal := o.FullIndexer[i]
		columnValueStr, ok := onIndexTotal.ColumnValue.(string)
		if ok {
			if strings.Contains(columnValueStr, likeStr) {
				for j := 0; j < len(onIndexTotal.Pointers); j++ {
					// 相同的话展开放入
					oneEle := &canCompareRanElement{
						Rank:    onIndexTotal.Pointers[j].Rank,
						Pointer: onIndexTotal.Pointers[j].Pointer,
					}
					searchResult = append(searchResult, oneEle)
					rankInResIndex[onIndexTotal.Pointers[j].Rank] = oneEle
				}
			}
		}
	}
	// 排序
	sort.Slice(searchResult, func(i, j int) bool {
		return searchResult[i].Rank > searchResult[j].Rank
	})
	for i := 0; i < len(searchResult); i++ {
		if i != 0 {
			searchResult[i-1].Next = searchResult[i]
			searchResult[i] = searchResult[i-1]
		}
	}
	return searchResult
}

type canCompareRanElement struct {
	Rank    int64
	Next    *canCompareRanElement
	Prove   *canCompareRanElement
	Pointer interface{}
}

type canCompareRanElements []*canCompareRanElement

// RetainAll 两个有序集合, c 交 b , 交集的集合在d 中
func (a canCompareRanElements) RetainAll(b canCompareRanElements) canCompareRanElements {
	var minRank, maxRank int64
	if a[0].Rank >= b[0].Rank {
		minRank = a[0].Rank
	} else if a[0].Rank <= b[0].Rank {
		minRank = b[0].Rank
	}
	aMaxlen := len(a) - 1
	bMaxlen := len(b) - 1
	if a[aMaxlen].Rank <= b[bMaxlen].Rank {
		maxRank = a[aMaxlen].Rank
	} else if a[aMaxlen].Rank >= b[bMaxlen].Rank {
		maxRank = b[bMaxlen].Rank
	}
	// 不存在
	if !(maxRank >= minRank) {
		return canCompareRanElements{}
	}
	var res = make([]*canCompareRanElement, 0)
	var aMin, aMax, bMin, bMax = a[0], a[aMaxlen], b[0], b[bMaxlen]
	if aMin.Rank >= bMin.Rank && aMax.Rank >= bMax.Rank {
		var bMap = make(map[int64]*canCompareRanElement)
		for current := bMin; current.Rank <= bMax.Rank; current = current.Next {
			bMap[current.Rank] = current
			if current.Next == nil {
				break
			}
		}
		for i := 0; i < len(a) && a[i].Rank >= bMin.Rank; i++ {
			if a[i].Rank > bMax.Rank {
				break
			}
			if bMap[a[i].Rank] != nil {
				res = append(res, bMap[a[i].Rank])
			}
		}
	} else if aMin.Rank < bMin.Rank && aMax.Rank > bMax.Rank {
		var aMap = make(map[int64]*canCompareRanElement)
		for current := aMin; current.Rank <= bMax.Rank; current = current.Next {
			aMap[current.Rank] = current
		}
		for i := 0; i < len(b) && b[i].Rank >= aMin.Rank; i++ {
			if aMap[b[i].Rank] != nil {
				res = append(res, aMap[b[i].Rank])
			}
		}
	} else if aMin.Rank > bMin.Rank && aMax.Rank < bMax.Rank {
		var bMap = make(map[int64]*canCompareRanElement)
		for current := bMax; current.Rank <= aMax.Rank; current = current.Next {
			bMap[current.Rank] = current
		}
		for i := 0; i < len(a) && a[i].Rank >= bMin.Rank; i++ {
			if bMap[a[i].Rank] != nil {
				res = append(res, bMap[a[i].Rank])
			}
		}
	} else if aMin.Rank <= bMin.Rank && aMax.Rank <= bMax.Rank {
		var aMap = make(map[int64]*canCompareRanElement)
		for current := aMin; current.Rank <= aMax.Rank; current = current.Next {
			aMap[current.Rank] = current
			if current.Next == nil {
				break
			}
		}
		for i := 0; i < len(b) && b[i].Rank >= bMin.Rank; i++ {
			if b[i].Rank > aMax.Rank {
				break
			}
			if aMap[b[i].Rank] != nil {
				res = append(res, aMap[b[i].Rank])
			}
		}
	}
	return res
}
