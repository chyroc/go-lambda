package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

type item struct {
	Name string
}

func Test_Map(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		var data interface{} = []*item{
			{Name: "a"},
			{Name: "b"},
		}

		obj, err := lambda.
			New(data).
			Map(func(idx int, v interface{}) interface{} {
				return v.(*item).Name
			}).
			Obj()
		as.Nil(err)
		as.Equal([]interface{}{"a", "b"}, obj)
	})

	t.Run("", func(t *testing.T) {
		obj, err := lambda.
			New(1).
			Map(func(idx int, v interface{}) interface{} {
				return v
			}).
			StringList()
		as.NotNil(err)
		as.Nil(obj)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})
}
