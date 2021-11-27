package rankdb

import "sync"

// SortedSetNode node of sorted set
type SortedSetNode struct {
	dict     map[interface{}]*sklNode
	skl      *skipList
	preWrite *skipList // 预写skl
	Lock     sync.Mutex
	Indexers *IndexerHash
}

func NewSortSet() *SortedSetNode {
	return &SortedSetNode{
		Indexers: NewIndexerHash(),
	}
}

// Add score 分值，value 保存的对象, indexers map 的key 保存的列名称,也就是value对象里面的某个字段的名称
// map 的 value 保存的value对象里面 key 对应的值。
// 这里让用户手动添加indexers 而不是利用反射是为了提高性能。减少反射所耗费的性能。
// TODO: 提供一个自动检查的方法。利用Tag标记。
func (s *SortedSetNode) Add(score CompareAble, value interface{}, indexers []*AddIndex) {
	s.Lock.Lock()
	insertEle := s.skl.sklInsert(score, value)
	insertRank := s.skl.sklGetRank(score, value)
	s.dict[value] = insertEle
	s.addIndexers(value, insertRank, indexers)
	s.Lock.Unlock()
}

func (s *SortedSetNode) addIndexers(value interface{}, valueRank int64, indexers []*AddIndex) {
	for i := 0; i < len(indexers); i++ {
		indexersKey, indexerVal := indexers[i].ColumnName, indexers[i].ColumnValue
		// 检查是否已经存在此字段的索引
		if oneIndexer := s.Indexers.IdxHash[indexersKey]; oneIndexer != nil {
			// 拉链法添加
			oneIndexer.ZipperAdd(indexerVal, value, valueRank)
		} else {
			// 新建对象添加
			newOneIndexer := NewOneIndexer()
			newOneIndexer.ZipperAdd(indexerVal, value, valueRank)
			s.Indexers.IdxHash[indexersKey] = newOneIndexer
		}
	}
}

// GetByIndex 传入搜索条件，查询指定key的结果。返回顺序按照此key的 score 排列
func (s *SortedSetNode) GetByIndex(selectIndexers []*SelectIndex, start, stop int64) []interface{} {
	var tempRes canCompareRanElements
	for i := 0; i < len(selectIndexers); i++ {
		var innerTempRes canCompareRanElements
		selectIndexer := selectIndexers[i]
		// 匹配索引
		columnName := selectIndexer.ColumnName
		// 检查传入字段是否是 索引队列
		if oneIndexer := s.Indexers.IdxHash[columnName]; oneIndexer != nil {
			// 判断索引值的类型，如果是字符串。支持模糊搜索
			if selectIndexer.ColumnType == STRING {
				if selectIndexer.Operation == LIKE { // 模糊搜索
					if searchValStr, ok := selectIndexer.SearchVal.(string); ok {
						// 获取模糊匹配的结果
						innerTempRes = oneIndexer.GetMemberLikeByStr(searchValStr)
					}
				}
			}
			// 等于操作,集合求交集
			if selectIndexer.Operation == EQ {
				if tempRes == nil {
					tempRes = innerTempRes
				} else {
					tempRes = tempRes.RetainAll(innerTempRes)
				}
			}
		}
	}
	res := make([]interface{}, len(tempRes))
	for i := 0; i < len(tempRes); i++ {
		res[i] = tempRes[i]
	}
	return res
}
