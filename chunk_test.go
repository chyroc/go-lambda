package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_chunk(t *testing.T) {
	as := assert.New(t)

	t.Run("chunk - success", func(t *testing.T) {
		resp, err := lambda.
			New([]int{1, 2, 3, 4, 5}).
			Chunk(2).
			ToIntListList()
		as.Nil(err)
		as.Equal([][]int{{1, 2}, {3, 4}, {5}}, resp)
	})

	t.Run("chunk - not-change-self", func(t *testing.T) {
		req := lambda.New([]int{1, 2, 3, 4, 5})
		req.Chunk(2)
		res, err := req.ToIntSlice()
		as.Nil(err)
		as.Equal([]int{1, 2, 3, 4, 5}, res)
	})

	t.Run("chunk - fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			Chunk(2).
			ToIntListList()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})

	t.Run("chunk - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			MapArray(func(idx int, obj interface{}) interface{} { return obj }).
			Chunk(2).
			ToIntListList()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})
}
