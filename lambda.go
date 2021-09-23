package lambda

type Object struct {
	err error
	obj interface{}
}

func New(obj interface{}) *Object {
	return &Object{obj: obj}
}
