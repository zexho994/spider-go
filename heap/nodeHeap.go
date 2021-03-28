package heap

import (
	"github.com/zouzhihao-994/spider-go/entity"
	"github.com/zouzhihao-994/spider-go/exception"
	"sync"
)

// stores nodes
type nodeHeap struct {
	space map[string]*BRoot
	count int32
}

var once sync.Once
var nh *nodeHeap

func NodeHeap() *nodeHeap {
	if nh == nil {
		once.Do(func() {
			nh = &nodeHeap{
				space: map[string]*BRoot{},
				count: 0,
			}
		})
	}
	return nh
}

// Construct a node
func (this nodeHeap) NewNode(name string, index string, attr map[string]string) (*entity.Node, error) {
	attribute := entity.NewAttribute(index, attr)
	n, err := entity.NewNode(name, attribute)
	if err != nil {
		return nil, err
	}
	this.count++
	//todo: add node to space

	return n, nil
}

// add a node to heap
func (this nodeHeap) addNode(name string) (*entity.Node, error) {
	node := this.space[name]
	if node == nil {
		return nil, exception.RuntimeException(exception.NodeNotFoundError)
	}

}
