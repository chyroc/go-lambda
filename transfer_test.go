package lambda_test

import (
	"strconv"
	"testing"

	"github.com/chyroc/go-assert"
	"github.com/chyroc/go-lambda"
)

func Test_Transfer(t *testing.T) {
	type res struct {
		item []*item
	}
	obj := lambda.New(res{item: []*item{
		{Name: "1"},
		{Name: "2"},
		{Name: "3"},
	}}).Transfer(func(obj interface{}) interface{} {
		return obj.(res).item
	}).Filter(func(idx int, obj interface{}) bool {
		i, _ := strconv.ParseInt(obj.(*item).Name, 10, 64)
		return i%2 == 0
	}).Map(func(idx int, v interface{}) interface{} {
		return v.(*item).Name
	}).Obj()

	assert.Equal(t, []interface{}{"2"}, obj)
}
