package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_ToList(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]string{"1", "2", "3"}).
				Filter(func(idx int, obj interface{}) bool {
					return obj.(string) != "2"
				}).
				Join("/")
			as.Nil(err)
			as.Equal("1/3", resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New(123).
				Join("/")
			as.Empty(resp)
			as.NotNil(err)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})
	})

	t.Run("", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]string{"1", "2", "3"}).
				Filter(func(idx int, obj interface{}) bool {
					return obj.(string) != "2"
				}).
				StringList()
			as.Nil(err)
			as.Equal([]string{"1", "3"}, resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New(123).
				StringList()
			as.Nil(resp)
			as.NotNil(err)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})
	})

	t.Run("", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			resp := []*item{}
			err := lambda.
				New([]string{"1", "2"}).
				Map(func(idx int, v interface{}) interface{} {
					return &item{Name: v.(string)}
				}).
				ToList(&resp)
			as.Nil(err)
			as.Len(resp, 2)
			as.Equal("1", resp[0].Name)
			as.Equal("2", resp[1].Name)
		})

		t.Run("", func(t *testing.T) {
			resp := []*item{}
			err := lambda.
				New([]string{"1", "2"}).
				Map(func(idx int, v interface{}) interface{} {
					return &item{Name: v.(string)}
				}).
				ToList(resp)
			as.NotNil(err)
			as.Equal("resp must be ptr", err.Error())
		})

		t.Run("", func(t *testing.T) {
			resp := []*item{}
			err := lambda.
				New(123).
				ToList(&resp)
			as.NotNil(err)
			as.Equal("int unsupport to array lambda operator", err.Error())
		})
	})

	t.Run("", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]rune{'1', '2', '3'}).
				Filter(func(idx int, obj interface{}) bool {
					return obj.(int32) != '2'
				}).
				String()
			as.Nil(err)
			as.Equal("13", resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]byte("123")).
				Filter(func(idx int, obj interface{}) bool {
					return obj.(uint8) != '2'
				}).
				String()
			as.Nil(err)
			as.Equal("13", resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New("13").
				String()
			as.Nil(err)
			as.Equal("13", resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New(123).
				String()
			as.Empty(resp)
			as.NotNil(err)
			as.Equal("int unsupport .String lambda operator", err.Error())
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]string{"1", "2", "3"}).
				Map(func(idx int, v interface{}) interface{} {
					return &item{Name: v.(string)}
				}).
				String()
			as.Empty(resp)
			as.NotNil(err)
			as.Equal("list of *lambda_test.item unsupport .String lambda operator", err.Error())
		})
	})
}
