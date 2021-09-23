package lambda_test

import (
	"strconv"
	"testing"

	"github.com/chyroc/go-assert"
	"github.com/chyroc/go-lambda"
)

func Test_Filter(t *testing.T) {
	as := assert.New(t)

	res := lambda.New([]string{"1", "2", "3"}).Filter(func(idx int, obj interface{}) bool {
		i, _ := strconv.ParseInt(obj.(string), 10, 64)
		return i%2 == 0
	}).Obj()
	as.Equal([]interface{}{"2"}, res)
}
