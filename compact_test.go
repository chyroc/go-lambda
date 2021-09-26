package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_compact(t *testing.T) {
	as := assert.New(t)

	t.Run("compact - success", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			resp, err := lambda.
				New([]int{0, 1, 2}).
				Compact().
				ToIntList()
			as.Nil(err)
			as.Equal([]int{1, 2}, resp)
		})
		t.Run("string", func(t *testing.T) {
			resp, err := lambda.
				New([]string{"", "x"}).
				Compact().
				ToStringList()
			as.Nil(err)
			as.Equal([]string{"x"}, resp)
		})
		t.Run("bool", func(t *testing.T) {
			resp, err := lambda.
				New([]bool{true, false, true}).
				Compact().
				ToBoolList()
			as.Nil(err)
			as.Equal([]bool{true, true}, resp)
		})
		t.Run("interface", func(t *testing.T) {
			resp, err := lambda.
				New([]interface{}{1, nilInterface}).
				Compact().
				ToInterfaceList()
			as.Nil(err)
			as.Equal([]interface{}{1}, resp)
		})
	})

	t.Run("compact - not-change-self", func(t *testing.T) {
		req := lambda.New([]int{0, 1, 2})
		req.Compact()
		res, err := req.ToIntList()
		as.Nil(err)
		as.Equal([]int{0, 1, 2}, res)
	})

	t.Run("compact - fail", func(t *testing.T) {
		_, err := lambda.
			New(234).
			Compact().
			ToIntList()
		as.NotNil(err)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})

	t.Run("compact - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(234).
			MapArray(func(idx int, obj interface{}) interface{} { return obj }).
			Compact().
			ToIntList()
		as.NotNil(err)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})
}
