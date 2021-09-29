package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/chyroc/go-ptr"
	"github.com/stretchr/testify/assert"
)

func Test_ToList(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]string{"1", "2", "3"}).
				FilterArray(func(idx int, obj interface{}) bool {
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
			as.Equal("123(int) can't convert to []string", err.Error())
		})
	})

	t.Run("", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			resp := []*item{}
			err := lambda.
				New([]string{"1", "2"}).
				MapArray(func(idx int, v interface{}) interface{} {
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
				MapArray(func(idx int, v interface{}) interface{} {
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
			as.Equal("123(int) can't convert to []interface", err.Error())
		})
		t.Run("", func(t *testing.T) {
			err := lambda.
				New([]string{"1"}).
				ToList(ptr.Int(1))
			as.NotNil(err)
			as.Equal("resp must be ptr of slice", err.Error())
		})
	})

	t.Run("", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]rune{'1', '2', '3'}).
				FilterArray(func(idx int, obj interface{}) bool {
					return obj.(int32) != '2'
				}).
				String()
			as.Nil(err)
			as.Equal("13", resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]byte("123")).
				FilterArray(func(idx int, obj interface{}) bool {
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
				MapArray(func(idx int, v interface{}) interface{} {
					return &item{Name: v.(string)}
				}).
				String()
			as.Empty(resp)
			as.NotNil(err)
			as.Equal("list of *lambda_test.item unsupport .String lambda operator", err.Error())
		})
	})
}
