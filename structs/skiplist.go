package structs

import (
	"math/rand"
	"strings"
)

type VALUE string

type SkipList struct {
	level        uint8
	length       uint16
	count        int64
	header, tail *SkipListNode
}

type SkipListNode struct {
	val      string
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
func (sl *SkipList) add(key string, val string, score float64) {
	sl.length++

	toInsert := &SkipListNode{
		level: make([]*SkipListLevel, sl.randomLevel()),
		score: score,
		val:   val,
	}

	pos := sl.findByScore(score)
	if pos.score == score {
		for pos.backward != nil && pos.backward.score == score && strings.Compare(val, pos.val) == -1 {
			pos = pos.backward
		}
		toInsert.backward = pos.backward
		toInsert.level[0].forward = pos
		if pos.backward != nil {
			pos.backward.level[0].forward = toInsert
		}
		pos.backward = toInsert
	} else {
		toInsert.level[0].forward = pos.level[0].forward
		toInsert.backward = pos
		if pos.level[0] != nil {
			pos.level[0].forward.backward = toInsert
		}
		pos.level[0].forward = toInsert
	}

	for i := len(toInsert.level) - 1; i > 0; i-- {
		if sl.header.level[i].forward == nil {
			sl.header.level[i].forward = toInsert
			continue
		}
		h := sl.header
		for h.level[i].forward != nil &&
			h.level[i].forward.score <= score &&
			strings.Compare(val, h.level[i].forward.val) == -1 {
			h = h.level[i].forward
		}

		if h.level[i].forward == nil {
			h.level[i].forward = pos
		} else {
			pos.level[i].forward = h.level[i].forward
			h.level[i].forward = pos
		}
	}

}

// remove a val from SkipList
func (SkipList) remove(key string, val VALUE) {

}

// range find SkipListNode by score
func (SkipList) rangeByScore(key string, s1, s2 float64) {

}

// find the SkipListNode by score, but it's not guaranteed to be the first one
func (sl *SkipList) findByScore(score float64) *SkipListNode {
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
	return node
}

// find the first node that it.score equals the score
func (sl *SkipList) firstForScore(score float64) *SkipListNode {
	r := sl.findByScore(score)
	if r == nil {
		return nil
	}
	for r.backward != nil {
		if r.backward.score != score {
			return r
		}
		r = r.backward
	}
	return r
}

// find the last node that it.score equals the score
func (sl *SkipList) lastForScore(score float64) *SkipListNode {
	r := sl.findByScore(score)
	if r == nil {
		return nil
	}

	for idx := len(r.level) - 1; idx >= 0; idx-- {
		if r.level[idx].forward.score > score {
			continue
		}
		r = r.level[idx].forward
		idx = len(r.level)
	}

	return r
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
