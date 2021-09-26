package lambda_test

import (
	"math"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_groupByArray(t *testing.T) {
	as := assert.New(t)

	t.Run("groupByArray - success", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			resp, err := lambda.
				New([]float64{0, 1.1, 2.2, 1.2, 2.3}).
				GroupByArray(func(idx int, obj interface{}) interface{} {
					return int(math.Ceil(obj.(float64))) // int(>= x)
				}).
				ToMapInt2Float64List()
			as.Nil(err)
			as.Equal(map[int][]float64{
				0: {0},
				2: {1.1, 1.2},
				3: {2.2, 2.3},
			}, resp)
		})
	})

	t.Run("groupByArray - fail", func(t *testing.T) {
		_, err := lambda.
			New(234).
			GroupByArray(func(idx int, obj interface{}) interface{} {
				return obj
			}).
			ToMapInt2Float64List()
		as.NotNil(err)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})

	t.Run("groupByArray - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(234).
			MapArray(func(idx int, obj interface{}) interface{} { return obj }).
			GroupByArray(func(idx int, obj interface{}) interface{} {
				return obj
			}).
			ToMapInt2Float64List()
		as.NotNil(err)
		as.Equal("int unsupport to array lambda operator", err.Error())
	})
}
