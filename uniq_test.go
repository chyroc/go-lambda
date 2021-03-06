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
				ToIntSlice()
			as.Nil(err)
			as.Equal([]int{0, 1, 2}, resp)
		})
	})

	t.Run("uniq - not-change-self", func(t *testing.T) {
		req := lambda.New([]int{0, 1, 2, 3, 2, 3})
		req.Uniq()
		res, err := req.ToIntSlice()
		as.Nil(err)
		as.Equal([]int{0, 1, 2, 3, 2, 3}, res)
	})

	t.Run("uniq - fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			Uniq().
			ToIntSlice()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})

	t.Run("uniq - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			MapList(func(idx int, obj interface{}) interface{} { return obj }).
			Uniq().
			ToIntSlice()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})
}
