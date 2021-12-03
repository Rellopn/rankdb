package rankdb

import (
	"fmt"
	"github.com/rellopn/rankdb/filter"
	"strings"
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
		Filter:        filter.And(filter.Eq{"a": "3"}),
		Level:         1,
		FatherBuilder: nil,
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
	bfs := queryBuilders.LevelBfs()
	for i, qd := range bfs {
		for _, q := range qd {
			fmt.Println("第 ", i, "层 [  ", q, "  ]")
		}
	}
}
func TestUnMarFilter(t *testing.T) {
	var nw = NewWriter()
	filter.Eq{"a": "1"}.
		And(filter.Eq{"b": "2"}).And(
		filter.Eq{"c": "3"}.And(
			filter.Eq{"d": 6}).And(filter.Eq{"f": 7}).And(filter.Eq{"g": 8})).WriteTo(nw)
	fmt.Println(nw.writer.String())

}

// BytesWriter implments Writer and save SQL in bytes.Buffer
type BytesWriter struct {
	writer *strings.Builder
	args   []interface{}
}

func NewWriter() *BytesWriter {
	w := &BytesWriter{
		writer: &strings.Builder{},
	}
	return w
}

// Write writes data to Writer
func (s *BytesWriter) Write(buf []byte) (int, error) {
	return s.writer.Write(buf)
}

// Append appends args to Writer
func (s *BytesWriter) Append(args ...interface{}) {
	s.args = append(s.args, args...)
}
