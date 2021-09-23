package lambda

import (
	"sync"
)

type Object struct {
	err error
	obj interface{}
	wg sync.WaitGroup
}

func New(obj interface{}) *Object {
	return &Object{obj: obj}
}
