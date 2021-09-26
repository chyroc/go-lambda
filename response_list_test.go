package lambda_test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_list(t *testing.T) {
	as := assert.New(t)

	t.Run("正确的测试 - string-list", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]string{"1", "2", "3"}).
				Filter(func(idx int, obj interface{}) bool {
					return obj.(string) != "2"
				}).
				ToStringList()
			as.Nil(err)
			as.Equal([]string{"1", "3"}, resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New(123).
				ToStringList()
			as.Nil(resp)
			as.NotNil(err)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})
	})

	t.Run("正确的测试 - int-list", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]int{1, 2, 3}).
				Filter(func(idx int, obj interface{}) bool {
					return obj.(int) != 2
				}).
				ToIntList()
			as.Nil(err)
			as.Equal([]int{1, 3}, resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New(123).
				ToStringList()
			as.Nil(resp)
			as.NotNil(err)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})
	})

	t.Run("出现了 err - string-list", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			obj, err := lambda.
				New(1).
				Array(func(idx int, v interface{}) interface{} {
					return v
				}).
				ToStringList()
			as.NotNil(err)
			as.Nil(obj)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})

		t.Run("", func(t *testing.T) {
			obj, err := lambda.
				New(1).
				ArrayAsync(func(idx int, v interface{}) interface{} {
					return v
				}).
				ToStringList()
			as.NotNil(err)
			as.Nil(obj)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})

		t.Run("常规 err", func(t *testing.T) {
			obj, err := lambda.
				New(1).
				ArrayAsyncWithErr(func(idx int, v interface{}) (interface{}, error) {
					return v, nil
				}).
				ToStringList()
			as.NotNil(err)
			as.Nil(obj)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})

		t.Run("WithErr 出现了 err", func(t *testing.T) {
			obj, err := lambda.
				New([]string{"1", "2", "3"}).
				ArrayAsyncWithErr(func(idx int, v interface{}) (interface{}, error) {
					if idx == 0 {
						return nil, fmt.Errorf("err")
					}
					return v, nil
				}).
				ToStringList()
			as.NotNil(err)
			as.Nil(obj)
			as.Equal("err", err.Error())
		})
	})
}
