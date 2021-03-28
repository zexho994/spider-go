package entity

import (
	"github.com/zouzhihao-994/spider-go/exception"
)

type Attribute struct {
	body map[string]string
}

func (this *Attribute) update(k, v string) (bool, error) {
	if this.body[k] == "" {
		return false, exception.RuntimeException(exception.AttributeNotExistException)
	}
	this.body[k] = v
	return true, nil
}
