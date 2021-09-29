package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	as := assert.New(t)

	t.Run("reverse - success", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			resp, err := lambda.
				New([]int{0, 1, 2}).
				Reverse().
				ToIntSlice()
			as.Nil(err)
			as.Equal([]int{2, 1, 0}, resp)
		})
	})

	t.Run("reverse - not-change-self", func(t *testing.T) {
		req := lambda.New([]int{0, 1, 2, 3})
		req.Reverse()
		res, err := req.ToIntSlice()
		as.Nil(err)
		as.Equal([]int{0, 1, 2, 3}, res)
	})

	t.Run("reverse - fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			Reverse().
			ToIntSlice()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})

	t.Run("reverse - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			MapList(func(idx int, obj interface{}) interface{} { return obj }).
			Reverse().
			ToIntSlice()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})
}
