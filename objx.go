package lambda

type Obj struct {
	err error
	obj interface{}
}

func New(obj interface{}) *Obj {
	return &Obj{obj: obj}
}
