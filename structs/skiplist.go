package structs

import "math/rand"

type VALUE string

type skipListMap map[string]*SkipList

type SkipList struct {
	level        uint8
	length       uint16
	count        int64
	header, tail *SkipListNode
}

type SkipListNode struct {
	val      VALUE
	score    float64
	level    []*SkipListLevel
	backward *SkipListNode
}

type SkipListLevel struct {
	forward *SkipListNode
	span    uint8
}

func initSkipList() *SkipList {
	return &SkipList{
		level:  0,
		length: 0,
		count:  0,
		header: &SkipListNode{level: make([]*SkipListLevel, 64)},
	}
}

// add or update the key&val to SkipList
func (SkipList) add(key string, val VALUE, score float64) {

}

// remove a val from SkipList
func (SkipList) remove(key string, val VALUE) {

}

// range find SkipListNode by score
func (SkipList) rangeByScore(key string, s1, s2 float64) {

}

// Generates a random number of level
func (sl *SkipList) randomLevel() uint8 {
	var i uint8 = 1
	for i <= sl.level {
		if uint8(rand.Intn(2)) == 1 {
			i++
		} else {
			return i
		}
	}
	sl.level++
	return sl.level
}
