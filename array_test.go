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

	t.Run("", func(t *testing.T) {
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
	})

	t.Run("", func(t *testing.T) {
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
	})

	t.Run("", func(t *testing.T) {
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
	})
}
