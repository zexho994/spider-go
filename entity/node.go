package entity

import "github.com/zouzhihao-994/spider-go/exception"

type Node struct {
	name         string
	attributes   *Attributes
	associated   map[string]*Relation
	beAssociated map[string]*Relation
}

func NewNode(name string, attributes *Attributes) (*Node, error) {
	if name == "" {
		return nil, exception.RuntimeException("name is blank")
	}
	if attributes == nil {
		return nil, exception.RuntimeException("attributes is nil")
	}
	return &Node{
		name:         name,
		attributes:   attributes,
		associated:   make(map[string]*Relation),
		beAssociated: make(map[string]*Relation),
	}, nil
}

func (this Node) Attribute() *Attributes {
	return this.attributes
}

func (this *Node) UpdateAttribute(k, v string) (bool, error) {
	return this.attributes.update(k, v)
}
