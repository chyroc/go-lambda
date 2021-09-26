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
				FilterArray(func(idx int, obj interface{}) bool {
					return obj.(int)%2 == 0
				}).
				ToIntList()
			as.Nil(err)
			as.Equal([]int{0, 2}, resp)
		})
	})

	t.Run("filterArray - fail", func(t *testing.T) {
		_, err := lambda.
			New(234).
			FilterArray(func(idx int, obj interface{}) bool {
				return true
			}).
			ToIntList()
		as.NotNil(err)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})

	t.Run("filterArray - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(234).
			MapArray(func(idx int, obj interface{}) interface{} { return obj }).
			FilterArray(func(idx int, obj interface{}) bool {
				return true
			}).
			ToIntList()
		as.NotNil(err)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})
}
