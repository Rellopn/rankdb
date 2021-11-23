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
