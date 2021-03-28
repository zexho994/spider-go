package entity

import "github.com/zouzhihao-994/spider-go/exception"

type node struct {
	name         string
	attributes   *Attribute
	associated   map[string]*Relation
	beAssociated map[string]*Relation
}

func Node(name string, attributes *Attribute) (*node, error) {
	if name == "" {
		return nil, exception.RuntimeException("name is blank")
	}
	if attributes == nil {
		return nil, exception.RuntimeException("attributes is nil")
	}
	return &node{
		name:         name,
		attributes:   attributes,
		associated:   make(map[string]*Relation),
		beAssociated: make(map[string]*Relation),
	}, nil
}

func (this *node) updateAttribute(k, v string) (bool, error) {
	return this.attributes.update(k, v)
}
