package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_uniq(t *testing.T) {
	as := assert.New(t)

	t.Run("uniq - success", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			resp, err := lambda.
				New([]int{0, 1, 0, 2, 1, 2}).
				Uniq().
				ToIntList()
			as.Nil(err)
			as.Equal([]int{0, 1, 2}, resp)
		})
	})

	t.Run("uniq - fail", func(t *testing.T) {
		_, err := lambda.
			New(234).
			Uniq().
			ToIntList()
		as.NotNil(err)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})

	t.Run("uniq - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(234).
			MapArray(func(idx int, obj interface{}) interface{} { return obj }).
			Uniq().
			ToIntList()
		as.NotNil(err)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})
}
