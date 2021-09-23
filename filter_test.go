package lambda_test

import (
	"strconv"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		res, err := lambda.
			New([]string{"1", "2", "3"}).
			Filter(func(idx int, obj interface{}) bool {
				i, _ := strconv.ParseInt(obj.(string), 10, 64)
				return i%2 == 0
			}).
			Obj()
		as.Nil(err)
		as.Equal([]interface{}{"2"}, res)
	})

	t.Run("", func(t *testing.T) {
		res, err := lambda.
			New("1/2/3/4/5").
			Filter(func(idx int, obj interface{}) bool {
				return obj.(rune) != '/'
			}).
			String()
		as.Nil(err)
		as.Equal("12345", res)
	})

	t.Run("", func(t *testing.T) {
		res, err := lambda.
			New(123).
			Filter(func(idx int, obj interface{}) bool {
				return true
			}).
			Obj()
		as.NotNil(err)
		as.Nil(res)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})
}
