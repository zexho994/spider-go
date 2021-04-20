package structs

type LEVEL uint8
type LENGTH uint16
type SCORE float64
type VALUE string

type skipListMap map[string]*SkipList

type SkipList struct {
	level  LEVEL
	length LENGTH
	header *SkipListNode
	tail   *SkipListNode
}

type SkipListNode struct {
	val   VALUE
	score SCORE
	level []struct {
		forward *SkipListNode
		span    LENGTH
	}
	bucket   []*SkipListNode
	backward *SkipListNode
}

// add or update the key&val to SkipList
func (skipListMap) add(key string, val VALUE, score SCORE) {

}

// remove a val from SkipList
func (skipListMap) remove(key string, val VALUE) {

}

// range find SkipListNode by score
func (skipListMap) rangeByScore(key string, s1, s2 SCORE) {

}
