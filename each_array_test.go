package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_eachArray(t *testing.T) {
	as := assert.New(t)

	t.Run("eachArray - success", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			res := []int{}
			err := lambda.New([]int{0, 1, 2}).
				EachArray(func(idx int, obj interface{}) {
					res = append(res, obj.(int))
				})
			as.Nil(err)
			as.Equal([]int{0, 1, 2}, res)
		})
	})

	t.Run("eachArray - fail", func(t *testing.T) {
		err := lambda.
			New(123).
			EachArray(func(idx int, obj interface{}) {
			})
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})

	t.Run("compact - pre-fail", func(t *testing.T) {
		err := lambda.
			New(123).
			MapArray(func(idx int, obj interface{}) interface{} { return obj }).
			EachArray(func(idx int, obj interface{}) {
			})
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})
}
