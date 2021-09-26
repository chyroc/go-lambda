package lambda

import (
	"fmt"
	"sync"
)

type Object struct {
	err error
	obj interface{}
	wg  *sync.WaitGroup
}

func New(obj interface{}) *Object {
	return &Object{obj: obj, wg: new(sync.WaitGroup)}
}

var ErrBreak = fmt.Errorf("break range")
