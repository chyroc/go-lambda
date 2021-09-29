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

	t.Run("array - success", func(t *testing.T) {
		t.Run("mapArray", func(t *testing.T) {
			var data interface{} = []*item{{Name: "a"}, {Name: "b"}}

			resp, err := lambda.New(data).
				MapArray(func(idx int, v interface{}) interface{} {
					return v.(*item).Name
				}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"a", "b"}, resp)
		})

		t.Run("mapArrayAsync", func(t *testing.T) {
			var data interface{} = []*item{{Name: "a"}, {Name: "b"}}

			resp, err := lambda.New(data).
				MapArrayAsync(func(idx int, v interface{}) interface{} {
					return v.(*item).Name
				}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"a", "b"}, resp)
		})

		t.Run("arrayAsyncWithErr", func(t *testing.T) {
			var data interface{} = []*item{{Name: "a"}, {Name: "b"}}

			resp, err := lambda.New(data).
				MapArrayAsyncWithErr(func(idx int, v interface{}) (interface{}, error) {
					return v.(*item).Name, nil
				}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"a", "b"}, resp)
		})
	})

	t.Run("array - fail", func(t *testing.T) {
		t.Run("array", func(t *testing.T) {
			t.Run("pre-fail", func(t *testing.T) {
				_, err := lambda.New(123).
					MapArray(func(idx int, obj interface{}) interface{} {
						return obj
					}).
					MapArray(func(idx int, v interface{}) interface{} {
						return v.(*item).Name
					}).ToStringSlice()
				as.NotNil(err)
				as.Equal("123(int) can't convert to []interface", err.Error())
			})

			t.Run("fail", func(t *testing.T) {
				_, err := lambda.New(123).
					MapArray(func(idx int, v interface{}) interface{} {
						return v.(*item).Name
					}).ToStringSlice()
				as.NotNil(err)
				as.Equal("123(int) can't convert to []interface", err.Error())
			})
		})

		t.Run("array-async", func(t *testing.T) {
			t.Run("pre-fail", func(t *testing.T) {
				_, err := lambda.New(123).
					MapArray(func(idx int, obj interface{}) interface{} {
						return obj
					}).
					MapArrayAsync(func(idx int, v interface{}) interface{} {
						return v.(*item).Name
					}).ToStringSlice()
				as.NotNil(err)
				as.Equal("123(int) can't convert to []interface", err.Error())
			})

			t.Run("fail", func(t *testing.T) {
				_, err := lambda.New(123).
					MapArrayAsync(func(idx int, v interface{}) interface{} {
						return v.(*item).Name
					}).ToStringSlice()
				as.NotNil(err)
				as.Equal("123(int) can't convert to []interface", err.Error())
			})
		})

		t.Run("array-async-err", func(t *testing.T) {
			t.Run("pre-fail", func(t *testing.T) {
				_, err := lambda.New(123).
					MapArray(func(idx int, obj interface{}) interface{} {
						return obj
					}).
					MapArrayAsyncWithErr(func(idx int, v interface{}) (interface{}, error) {
						return v.(*item).Name, nil
					}).ToStringSlice()
				as.NotNil(err)
				as.Equal("123(int) can't convert to []interface", err.Error())
			})

			t.Run("fail", func(t *testing.T) {
				_, err := lambda.New(123).
					MapArrayAsyncWithErr(func(idx int, v interface{}) (interface{}, error) {
						return v.(*item).Name, nil
					}).ToStringSlice()
				as.NotNil(err)
				as.Equal("123(int) can't convert to []interface", err.Error())
			})
		})
	})

	t.Run("async - success", func(t *testing.T) {
		t.Run("array - dur > 2", func(t *testing.T) {
			start := time.Now()
			resp, err := lambda.New([]string{"1", "2", "3"}).
				MapArray(func(idx int, v interface{}) interface{} {
					time.Sleep(time.Second)
					return v
				}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"1", "2", "3"}, resp)
			dur := time.Now().Sub(start).Seconds()
			as.True(dur > 2)
		})

		t.Run("async - duc < 2", func(t *testing.T) {
			start := time.Now()
			resp, err := lambda.New([]string{"1", "2", "3"}).
				MapArrayAsync(func(idx int, v interface{}) interface{} {
					time.Sleep(time.Second)
					return v
				}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"1", "2", "3"}, resp)
			dur := time.Now().Sub(start).Seconds()
			as.True(dur < 2)
		})

		t.Run("async-err - dur < 2", func(t *testing.T) {
			start := time.Now()
			resp, err := lambda.New([]string{"1", "2", "3"}).
				MapArrayAsyncWithErr(func(idx int, v interface{}) (interface{}, error) {
					time.Sleep(time.Second)
					return v, nil
				}).ToStringSlice()
			as.Nil(err)
			as.Equal([]string{"1", "2", "3"}, resp)
			dur := time.Now().Sub(start).Seconds()
			as.True(dur < 2)
		})
	})
}
