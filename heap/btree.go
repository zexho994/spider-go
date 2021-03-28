package heap

import "github.com/zouzhihao-994/spider-go/entity"

// b+ tree, stores the node and sort by index of attribute
type BRoot struct {
}

type Bnode interface {
	Val()
}

type bTreeNode struct {
	val   *entity.Node
	left  *Bnode
	right *Bnode
}

func (this bTreeNode) Val() *entity.Node {
	return this.val
}

func (this bTreeNode) insert(node *entity.Node) {
	node.Attribute()
}
