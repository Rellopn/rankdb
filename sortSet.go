package rankdb

import (
	"sync"
)

type SortSetRecord map[string]*SortSet

var SortSetRecords = make(map[string]*SortSet)

type SortSet struct {
	Lock     sync.Mutex
	Indexers *IndexerHash
	SkipList *SkipList
}

func NewSortSet() *SortSet {
	return &SortSet{
		Indexers: NewIndexerHash(),
		SkipList: NewSkl(),
	}
}

func (s *SortSet) Add(score CompareAble, value interface{}, indexers map[string]string) {
	s.Lock.Lock()
	s.SkipList.Insert(score, value)
	s.addIndexers(value, indexers)
	s.Lock.Unlock()
}

func (s *SortSet) addIndexers(value interface{}, indexers map[string]string) {
	for indexersKey, IndexerVal := range indexers {
		if s.Indexers.IdxHash[indexersKey] != nil {
			oneIndexerFullIndexer := s.Indexers.IdxHash[indexersKey].FullIndexer
			oneIndexerFullIndexer = append(oneIndexerFullIndexer, &Indexer{ColumnName: IndexerVal, Pointer: value})
		} else {
			s.Indexers.IdxHash[indexersKey] = &OneIndexer{FullIndexer: []*Indexer{
				{ColumnName: IndexerVal, Pointer: value},
			}}
		}
	}
}

// ZCard 获取有序集合的成员数
func ZCard(key string) int {
	record := SortSetRecords[key]
	if record == nil {
		return 0
	}
	return record.SkipList.Nodes[0].EleNum
}
