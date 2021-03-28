package entity

import (
	"github.com/zouzhihao-994/spider-go/exception"
)

type Attributes struct {
	index string
	data  map[string]string
}

func NewAttribute(index string, attr map[string]string) *Attributes {
	return &Attributes{
		index: index,
		data:  attr,
	}
}

func (this Attributes) Index() string {
	return this.index
}

func (this *Attributes) get(k string) string {
	return this.data[k]
}

func (this *Attributes) update(k, v string) (bool, error) {
	if this.data[k] == "" {
		return false, exception.RuntimeException(exception.AttributeNotExistError)
	}
	this.data[k] = v
	return true, nil
}
