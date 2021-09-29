package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_indexOf(t *testing.T) {
	as := assert.New(t)

	t.Run("indexOf - success", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			resp, err := lambda.
				New([]int{0, 1, 2}).
				IndexOf(int(1)).
				ToInt()
			as.Nil(err)
			as.Equal(int(1), resp)
		})

		t.Run("not-found", func(t *testing.T) {
			resp, err := lambda.
				New([]int{0, 1, 2}).
				IndexOf(int(3)).
				ToInt()
			as.Nil(err)
			as.Equal(int(-1), resp)
		})
	})

	t.Run("indexOf - not-change-self", func(t *testing.T) {
		req := lambda.New([]int{0, 1, 2, 3})
		req.IndexOf(1)
		res, err := req.ToIntSlice()
		as.Nil(err)
		as.Equal([]int{0, 1, 2, 3}, res)
	})

	t.Run("indexOf - fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			IndexOf(1).
			ToInt()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})

	t.Run("indexOf - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			MapArray(func(idx int, obj interface{}) interface{} { return obj }).
			IndexOf(1).
			ToInt()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})
}
