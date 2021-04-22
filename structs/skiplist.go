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

// find the SkipListNode by score, but it's not guaranteed to be the first one
func (sl *SkipList) findByScore(val string, score float64) *SkipListNode {
	node := sl.header
	for l := len(node.level); l >= 0; l-- {
		if node.level[l].forward == nil || node.level[l].forward.score > score {
			continue
		}
		node = node.level[l].forward
		if node.score == score {
			return node
		}
		l = len(node.level)
	}
	return nil
}

func findK(sl *SkipListNode, score float64, size int64) {

}

// Generates a random number of level
func (sl *SkipList) randomLevel() uint8 {
	for i := uint8(1); i <= sl.level; i++ {
		if uint8(rand.Intn(2)) == 0 {
			return i
		}
	}
	if sl.level < 64 {
		sl.level++
	}
	return sl.level
}
