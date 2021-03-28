package entity

type node struct {
	name         string
	attributes   *Attribute
	associated   map[string]*Relation
	beAssociated map[string]*Relation
}
