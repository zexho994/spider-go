package _struct

var skiplistMap = make(map[string]*SkipList, 16)

type SkipList struct {
	level  uint8
	length uint16
	header *SkipListHeader
	tail   *SkipListNode
}

type SkipListHeader struct {
	data []*SkipListNode
}

type SkipListNode struct {
	val   string
	score uint16
	next  *SkipListNode
}

// add or update the key&val to SkipList
func add(key, val string, score uint16) {
	if skiplistMap[key] == nil {
		skiplistMap[key] = &SkipList{
			length: 0,
			level:  1 << 7,
			header: &SkipListHeader{
				data: make([]*SkipListNode, 16),
			},
			tail: nil,
		}
	} else {
	}

}

// remove a val from SkipList
func remove(key, val string) {

}

// range find SkipListNode by score
func (SkipList) rangeByScore(key, s1, s2 uint16) {

}
