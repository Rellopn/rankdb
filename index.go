package rankdb

// Indexer 实现了从sort set value中提取字段作为索引
type Indexer struct {
	// 字段的名称
	ColumnName string
	// 字段所对应的 结构体的指针
	Pointer interface{}
}

// OneIndexer 一条索引的全部记录
type OneIndexer struct {
	IndexName   string
	FullIndexer []*Indexer
}
