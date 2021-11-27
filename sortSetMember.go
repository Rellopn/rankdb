package rankdb

type AddIndex struct {
	ColumnName  string
	ColumnValue interface{}
}

type M struct {
	Score      CompareAble
	Member     interface{}
	AddIndexes []*AddIndex
}

const (
	EQ = iota
	LIKE
)
const (
	INT = iota
	FLOAT
	STRING
	TIME
)

// SelectIndex 搜索索引条件
type SelectIndex struct {
	ColumnName string      // 列名称
	Operation  int         // 操作 EQ LIKE ...
	SearchVal  interface{} // 查找的列的值
	ColumnType int         // 列的值的类型 支持 INT FLOAT STRING TIME
}

func NewSelectIndex() *SelectIndex {
	return &SelectIndex{}
}
