package lambda_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

type item struct {
	Name string
}

func Test_Map(t *testing.T) {
	as := assert.New(t)

	t.Run("正常处理", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			var data interface{} = []*item{
				{Name: "a"},
				{Name: "b"},
			}

			obj, err := lambda.
				New(data).
				Array(func(idx int, v interface{}) interface{} {
					return v.(*item).Name
				}).
				Obj()
			as.Nil(err)
			as.Equal([]interface{}{"a", "b"}, obj)
		})

		t.Run("", func(t *testing.T) {
			var data interface{} = []*item{
				{Name: "a"},
				{Name: "b"},
			}

			obj, err := lambda.
				New(data).
				ArrayAsync(func(idx int, v interface{}) interface{} {
					return v.(*item).Name
				}).
				Obj()
			as.Nil(err)
			as.Equal([]interface{}{"a", "b"}, obj)
		})

		t.Run("", func(t *testing.T) {
			var data interface{} = []*item{
				{Name: "a"},
				{Name: "b"},
			}

			obj, err := lambda.
				New(data).
				ArrayAsyncWithErr(func(idx int, v interface{}) (interface{}, error) {
					return v.(*item).Name, nil
				}).
				Obj()
			as.Nil(err)
			as.Equal([]interface{}{"a", "b"}, obj)
		})
	})

	t.Run("真的异步了，sleep", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			start := time.Now()
			obj, err := lambda.
				New([]string{"1", "2", "3"}).
				Array(func(idx int, v interface{}) interface{} {
					time.Sleep(time.Second)
					return v
				}).
				Obj()
			as.Nil(err)
			as.Equal([]interface{}{"1", "2", "3"}, obj)
			dur := time.Now().Sub(start).Seconds()
			as.True(dur > 2)
		})

		t.Run("", func(t *testing.T) {
			start := time.Now()
			obj, err := lambda.
				New([]string{"1", "2", "3"}).
				ArrayAsync(func(idx int, v interface{}) interface{} {
					time.Sleep(time.Second)
					return v
				}).
				Obj()
			as.Nil(err)
			as.Equal([]interface{}{"1", "2", "3"}, obj)
			dur := time.Now().Sub(start).Seconds()
			as.True(dur < 2)
		})

		t.Run("", func(t *testing.T) {
			start := time.Now()
			obj, err := lambda.
				New([]string{"1", "2", "3"}).
				ArrayAsyncWithErr(func(idx int, v interface{}) (interface{}, error) {
					time.Sleep(time.Second)
					return v, nil
				}).
				Obj()
			as.Nil(err)
			as.Equal([]interface{}{"1", "2", "3"}, obj)
			dur := time.Now().Sub(start).Seconds()
			as.True(dur < 2)
		})
	})

	t.Run("出现了 err", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			obj, err := lambda.
				New(1).
				Array(func(idx int, v interface{}) interface{} {
					return v
				}).
				StringList()
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
				StringList()
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
				StringList()
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
				StringList()
			as.NotNil(err)
			as.Nil(obj)
			as.Equal("err", err.Error())
		})
	})
}
