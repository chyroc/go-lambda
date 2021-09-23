package lambda_test

import (
	"strconv"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_Transfer(t *testing.T) {
	as := assert.New(t)

	type res struct {
		item []*item
	}
	obj, err := lambda.New(res{item: []*item{
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

	as.Equal([]interface{}{"2"}, obj)
	as.Nil(err)
}
