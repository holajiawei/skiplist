package skiplist

import (
	"math/rand"
	"sync"
)

type node struct {
	next  []*node
	prev  *node
	value SkipListItem
	mux   sync.Mutex
}

type SkipList struct {
	MaxLevels   int
	header      *node
	footer      *node
	length      int
	probability float64
	lock        sync.RWMutex
}

type SkipListItem interface {
	Less(b SkipListItem) bool
	Equals(b SkipListItem) bool
}

func New(probability float64, MaxLevels int) *SkipList {
	list := new(SkipList)
	list.probability = probability
	list.MaxLevels = MaxLevels
	return list
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (s *SkipList) effectiveMaxLevel() int {
	return maxInt(s.level(), s.MaxLevels)
}

func (s *SkipList) level() int {
	return len(s.header.next) - 1
}

func (s *SkipList) randomLevel() (n int) {
	for n = 0; n < s.MaxLevels && rand.Float64() < s.probability; n++ {
	}
	return
}
