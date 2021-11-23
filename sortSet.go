package rankdb

import (
	"sync"
)

// SortedSetNode node of sorted set
type SortedSetNode struct {
	dict     map[interface{}]*sklNode
	skl      *skipList
	Lock     sync.Mutex
	Indexers *IndexerHash
}

func NewSortSet() *SortedSetNode {
	return &SortedSetNode{
		Indexers: NewIndexerHash(),
		//skl: NewSkl(),
	}
}

// Add score 分值，value 保存的对象, indexers map 的key 保存的列名称,也就是value对象里面的某个字段的名称
// map 的 value 保存的value对象里面 key 对应的值。
// 这里让用户手动添加indexers 而不是利用反射是为了提高性能。减少反射所耗费的性能。
// TODO: 提供一个自动检查的方法。利用Tag标记。
func (s *SortedSetNode) Add(score CompareAble, value interface{}, indexers []*AddIndex) {
	s.Lock.Lock()
	insertEle := s.skl.sklInsert(score, value)
	s.dict[value] = insertEle
	s.addIndexers(value, indexers)
	s.Lock.Unlock()
}

func (s *SortedSetNode) addIndexers(value interface{}, indexers []*AddIndex) {
	for i := 0; i < len(indexers); i++ {
		indexersKey, indexerVal := indexers[i].ColumnName, indexers[i].ColumnValue
		// 检查是否已经存在此字段的索引
		if oneIndexer := s.Indexers.IdxHash[indexersKey]; oneIndexer != nil {
			// 拉链法添加
			oneIndexer.ZipperAdd(indexerVal, value)
		} else {
			// 新建对象添加
			newOneIndexer := NewOneIndexer()
			newOneIndexer.ZipperAdd(indexerVal, value)
			s.Indexers.IdxHash[indexersKey] = newOneIndexer
		}
	}
}
