package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_filterArray(t *testing.T) {
	as := assert.New(t)

	t.Run("filterArray - success", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			resp, err := lambda.
				New([]int{0, 1, 2, 3}).
				FilterList(func(idx int, obj interface{}) bool {
					return obj.(int)%2 == 0
				}).
				ToIntSlice()
			as.Nil(err)
			as.Equal([]int{0, 2}, resp)
		})
	})

	t.Run("filterArray - fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			FilterList(func(idx int, obj interface{}) bool {
				return true
			}).
			ToIntSlice()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})

	t.Run("filterArray - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			MapList(func(idx int, obj interface{}) interface{} { return obj }).
			FilterList(func(idx int, obj interface{}) bool {
				return true
			}).
			ToIntSlice()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})

	t.Run("filterArray - not-change-self", func(t *testing.T) {
		req := lambda.New([]int{0, 1, 2})
		req.FilterList(func(idx int, obj interface{}) bool {
			return obj.(int)%2 == 0
		})
		res, err := req.ToIntSlice()
		as.Nil(err)
		as.Equal([]int{0, 1, 2}, res)
	})
}
