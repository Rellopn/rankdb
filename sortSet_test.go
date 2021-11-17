package rankdb

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type RaceRankCompareAble struct {
	FinishPass int     // 1
	Speed      float64 // 2
	Distance   int     // 3
	totalTime  float64 // 4
}

func (r RaceRankCompareAble) Compare(a CompareAble) (int, error) {
	ad, ok := a.(RaceRankCompareAble)
	if !ok {
		aRealKind := reflect.TypeOf(a).Name()
		return 0, errors.New(fmt.Sprintf("Not except type, except type RaceRank, but got [%s]", aRealKind))
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
	sortSet := NewSortSet("test1")
	sortSet.Add("test1", &RaceRankCompareAble{
		FinishPass: 1,
		Speed:      1515.12357,
		Distance:   556879,
		totalTime:  367.54692,
	}, &RaceRankResult{
		Name:      "mark",
		Age:       23,
		ShoeBrand: "Nk",
		Speed:     1515.12357,
	}, map[string]string{"name": "mark"})
	sortSet.SkipList.Print()
}
