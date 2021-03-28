package entity

import (
	"github.com/zouzhihao-994/spider-go/exception"
)

type Attribute struct {
	body map[string]string
}

func NewAttribute(attr map[string]string) *Attribute {
	return &Attribute{
		body: attr,
	}
}

func (this *Attribute) update(k, v string) (bool, error) {
	if this.body[k] == "" {
		return false, exception.RuntimeException(exception.AttributeNotExistError)
	}
	this.body[k] = v
	return true, nil
}
