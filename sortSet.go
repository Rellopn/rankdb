package rankdb

import "sync"

type SortSet struct {
	Lock     sync.Mutex
	Indexers map[string]*OneIndexer
	SkipList *SkipList
}

func NewSortSet(key string) *SortSet {
	return &SortSet{
		Indexers: make(map[string]*OneIndexer),
		SkipList: NewSkl(key),
	}
}

func (s *SortSet) Add(key string, score CompareAble, value interface{}, indexers map[string]string) {
	s.Lock.Lock()
	s.SkipList.Insert(key, score, value)
	s.addIndexers(value, indexers)
	s.Lock.Unlock()
}

func (s *SortSet) addIndexers(value interface{}, indexers map[string]string) {
	for indexersKey, IndexerVal := range indexers {
		if s.Indexers[indexersKey] != nil {
			oneIndexerFullIndexer := s.Indexers[indexersKey].FullIndexer
			oneIndexerFullIndexer = append(oneIndexerFullIndexer, &Indexer{ColumnName: IndexerVal, Pointer: value})
		} else {
			s.Indexers[indexersKey] = &OneIndexer{IndexName: indexersKey, FullIndexer: []*Indexer{
				{ColumnName: IndexerVal, Pointer: value},
			}}
		}
	}
}
