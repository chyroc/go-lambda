package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_list(t *testing.T) {
	as := assert.New(t)

	t.Run("to-string", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			res, err := lambda.New(
				[]string{"1", "2"},
			).ToStringList()
			as.Nil(err)
			as.Equal([]string{"1", "2"}, res)
		})

		t.Run("fail", func(t *testing.T) {
			_, err := lambda.New(
				123,
			).ToStringList()
			as.NotNil(err)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})
	})

	t.Run("to-int", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			res, err := lambda.New(
				[]int{1, 2},
			).ToIntList()
			as.Nil(err)
			as.Equal([]int{1, 2}, res)
		})

		t.Run("fail", func(t *testing.T) {
			_, err := lambda.New(
				123,
			).ToIntList()
			as.NotNil(err)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})
	})

	t.Run("to-bool", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			res, err := lambda.New(
				[]bool{false, true},
			).ToBoolList()
			as.Nil(err)
			as.Equal([]bool{false, true}, res)
		})

		t.Run("fail", func(t *testing.T) {
			_, err := lambda.New(
				123,
			).ToBoolList()
			as.NotNil(err)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})
	})

	t.Run("to-interface{}", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			res, err := lambda.New(
				[]interface{}{1, 2},
			).ToInterfaceList()
			as.Nil(err)
			as.Equal([]interface{}{1, 2}, res)
		})

		t.Run("fail", func(t *testing.T) {
			_, err := lambda.New(
				123,
			).ToInterfaceList()
			as.NotNil(err)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})
	})
}
