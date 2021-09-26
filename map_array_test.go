package lambda_test

import (
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
				MapArray(func(idx int, v interface{}) interface{} {
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
				MapArray(func(idx int, v interface{}) interface{} {
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
}
