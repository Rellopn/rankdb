package rankdb

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type RaceRankCompareAble struct {
	FinishPass int     // 1
	Speed      float64 // 2
	Distance   int     // 3
	totalTime  float64 // 4
}

func (r RaceRankCompareAble) CompareTo(a CompareAble) (int, error) {
	ad, ok := a.(*RaceRankCompareAble)
	if !ok {
		aRealKind := reflect.TypeOf(a).Elem().Name()
		return 0, errors.New(fmt.Sprintf("Not except type, except type *RaceRankCompareAble, but got [%s]", aRealKind))
	}
	if r.FinishPass > ad.FinishPass {
		return 1, nil
	} else if r.FinishPass < ad.FinishPass {
		return -1, nil
	} else {
		if r.Speed > ad.Speed {
			return 1, nil
		} else if r.Speed < ad.Speed {
			return -1, nil
		} else {
			if r.Distance > ad.Distance {
				return 1, nil
			} else if r.Distance < ad.Distance {
				return -1, nil
			} else {
				if r.totalTime > ad.totalTime {
					return 1, nil
				} else if r.totalTime < ad.totalTime {
					return -1, nil
				} else {
					return 0, nil
				}
			}
		}
	}
}

type RaceRankResult struct {
	Rank      int
	Name      string
	Age       int
	ShoeBrand string
	Speed     float64
}

func TestSortSet_Add(t *testing.T) {
	sortSet := NewSortSet()
	sortSet.Add(&RaceRankCompareAble{
		FinishPass: 1,
		Speed:      1515.12357,
		Distance:   556879,
		totalTime:  367.54692,
	}, &RaceRankResult{
		Name:      "mark",
		Age:       23,
		ShoeBrand: "Nk",
		Speed:     1515.12357,
	}, []*AddIndex{
		{"name", "mark"},
	})
}

func TestZCard(t *testing.T) {
	sortSet := NewSortSet()
	for i := 0; i < 10000; i++ {
		sortSet.Add(&RaceRankCompareAble{
			FinishPass: 1,
			Speed:      1515.12357,
			Distance:   556879,
			totalTime:  367.54692,
		}, &RaceRankResult{
			Name:      "mark",
			Age:       23,
			ShoeBrand: "Nk",
			Speed:     1515.12357,
		}, []*AddIndex{
			{"name", "mark"},
		})
	}
}
func BenchmarkZCard(b *testing.B) {
	sortSet := NewSortSet()
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sortSet.Add(&RaceRankCompareAble{
				FinishPass: 1,
				Speed:      1515.12357,
				Distance:   556879,
				totalTime:  367.54692,
			}, &RaceRankResult{
				Name:      "mark",
				Age:       23,
				ShoeBrand: "Nk",
				Speed:     1515.12357,
			}, []*AddIndex{
				{"name", "mark"},
			})
		}()
	}
	wg.Wait()
	//keyNum := ZCard("t1")
	//fmt.Println("except 101, got ", keyNum)
}
