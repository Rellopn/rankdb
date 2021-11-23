package rankdb

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
	Pointers []interface{}
}

func (i *Indexer) appendPointers(pointer interface{}) {
	i.Pointers = append(i.Pointers, pointer)
}

// OneIndexer 一条索引的全部记录
type OneIndexer struct {
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

func (o *OneIndexer) ZipperAdd(columnValues, pointer interface{}) {
	if zipperIndex := o.StoreColumnValueExit(columnValues); zipperIndex != 0 {
		// 存在的话添加数组元素
		o.FullIndexer[zipperIndex-1].appendPointers(pointer)
	} else {
		o.storeColumnValues[columnValues] = 1
		o.FullIndexer = append(o.FullIndexer,
			&Indexer{ColumnValue: columnValues, Pointers: []interface{}{pointer}})
	}
}

func (o *OneIndexer) getFullIndexerLen() int {
	return len(o.FullIndexer)
}
