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
				FilterList(func(idx int, obj interface{}) bool {
					return obj.(string) != "2"
				}).
				ToJoin("/")
			as.Nil(err)
			as.Equal("1/3", resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New(123).
				ToJoin("/")
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
				MapList(func(idx int, v interface{}) interface{} {
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
				MapList(func(idx int, v interface{}) interface{} {
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
				New("13").
				ToString()
			as.Nil(err)
			as.Equal("13", resp)
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New(123).
				ToString()
			as.Empty(resp)
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert to")
		})

		t.Run("", func(t *testing.T) {
			resp, err := lambda.
				New([]string{"1", "2", "3"}).
				MapList(func(idx int, v interface{}) interface{} {
					return &item{Name: v.(string)}
				}).
				ToString()
			as.Empty(resp)
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert to")
		})
	})
}
