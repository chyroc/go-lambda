package lambda_test

import (
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_Flatten(t *testing.T) {
	as := assert.New(t)

	t.Run("出错", func(t *testing.T) {
		as.NotNil(lambda.New(1).Flatten().Error())
	})

	t.Run("没问题", func(t *testing.T) {
		// 1 2 3 => [1 1] [2 2] [3 3] => [1 1 2 2 3 3]
		res, err := lambda.New([]string{"1", "2", "3"}).Array(func(idx int, obj interface{}) interface{} {
			x := obj.(string)
			return []string{x, x}
		}).Flatten().Join("")
		as.Nil(err)
		as.Equal("112233", res)
	})

	t.Run("没问题", func(t *testing.T) {
		// 1 2 3 => [1 1] 2 3 => [1 1 2 3]
		res, err := lambda.New([]string{"1", "2", "3"}).Array(func(idx int, obj interface{}) interface{} {
			x := obj.(string)
			if x == "1" {
				return []string{x, x}
			}
			return x
		}).Flatten().Join("")
		as.Nil(err)
		as.Equal("1123", res)
	})
}
