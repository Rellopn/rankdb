package rankdb

import (
	"fmt"
	"github.com/rellopn/rankdb/filter"
	"testing"
)

func TestMatch(t *testing.T) {
	_ = QueryBuilder{
		Filter: filter.And(filter.Eq{"a": "3"}),
		InnerBuilders: []*QueryBuilder{
			{
				Filter: filter.And(filter.Eq{"b": "4"}),
			},
			{
				Filter: filter.Or(filter.Eq{"c": "5"}),
				InnerBuilders: []*QueryBuilder{
					{
						Filter: filter.And(filter.Eq{"d": "6"}),
					},
					{
						Filter: filter.Or(filter.Eq{"c": "5"}),
					},
				},
			},
		},
	}
}

func TestQueryBuilder_LevelBfs(t *testing.T) {
	qb := []*QueryBuilder{{
		Filter: filter.And(filter.Eq{"a": "3"}),
		Level:  1,
		InnerBuilders: []*QueryBuilder{
			{
				Filter: filter.And(filter.Eq{"b": "4"}),
				Level:  2,
			},
			{
				Filter: filter.Or(filter.Eq{"c": "5"}),
				Level:  2,
				InnerBuilders: []*QueryBuilder{
					{
						Filter: filter.And(filter.Eq{"d": "6"}),
						Level:  3,
						InnerBuilders: []*QueryBuilder{
							{
								Filter: filter.And(filter.Eq{"f": "7"}),
								Level:  4,
							},
							{
								Filter: filter.Or(filter.Eq{"G": "8"}),
								Level:  4,
							},
						},
					},
					{
						Filter: filter.Or(filter.Eq{"e": "6"}),
						Level:  3,
						InnerBuilders: []*QueryBuilder{
							{
								Filter: filter.And(filter.Eq{"h": "9"}),
								Level:  4,
							},
							{
								Filter: filter.Or(filter.Eq{"i": "10"}),
								Level:  4,
							},
						},
					},
				},
			},
		},
	},
	}
	queryBuilders := QueryBuilders{MaxLevel: 4, QueryBuilders: qb}
	for i, qd := range queryBuilders.LevelBfs() {
		for _, q := range qd {
			fmt.Println("第 ", i, "层 [  ", q, "  ]")
		}
	}
}
