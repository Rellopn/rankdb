package rankdb

import (
	"github.com/rellopn/rankdb/filter"
)

const (
	OperationAnd = iota
	OperationOr
)

type Operation struct {
	ConnectionSymbol int // 连接符号 AND, OR
}

// QueryBuilder 构建一颗查询的多叉树
type QueryBuilder struct {
	filter.Filter
	Level         int // 层级
	InnerBuilders []*QueryBuilder
	FatherBuilder *QueryBuilder // 上级节点
}

type QueryBuilders struct {
	QueryBuilders []*QueryBuilder
	MaxLevel      int // 当前树有几层
}

func (q *QueryBuilders) LevelBfs() [][]*QueryBuilder {
	res := make([][]*QueryBuilder, q.MaxLevel)
	queues := make([]*Queue, q.MaxLevel)
	q.levelBfs(q.QueryBuilders, &queues)
	for i := 0; i < len(queues); i++ {
		l := queues[i].Length()
		res[i] = make([]*QueryBuilder, l)
		for j := 0; j < l; j++ {
			res[i][j] = queues[i].Remove()
		}
	}
	return res
}
func (q *QueryBuilders) levelBfs(qbs []*QueryBuilder, queues *[]*Queue) {
	for i := 0; i < len(qbs); i++ {
		getReal := *queues
		if getReal[qbs[i].Level-1] == nil {
			getReal[qbs[i].Level-1] = NewQueue()
		}
		getReal[qbs[i].Level-1].Add(qbs[i])
		queues = &getReal
		if qbs[i].InnerBuilders != nil || len(qbs[i].InnerBuilders) != 0 {
			q.levelBfs(qbs[i].InnerBuilders, queues)
		}
	}
}
