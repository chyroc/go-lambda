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
				ToMapInt2Float64Slice()
			as.Nil(err)
			as.Equal(map[int][]float64{
				0: {0},
				2: {1.1, 1.2},
				3: {2.2, 2.3},
			}, resp)
		})
	})

	t.Run("groupByArray - not-change-self", func(t *testing.T) {
		req := lambda.New([]int{0, 1, 2, 3})
		req.GroupByArray(func(idx int, obj interface{}) interface{} {
			if obj.(int) > 1 {
				return 2
			} else {
				return 1
			}
		})
		res, err := req.ToIntSlice()
		as.Nil(err)
		as.Equal([]int{0, 1, 2, 3}, res)
	})

	t.Run("groupByArray - fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			GroupByArray(func(idx int, obj interface{}) interface{} {
				return obj
			}).
			ToMapInt2Float64Slice()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})

	t.Run("groupByArray - pre-fail", func(t *testing.T) {
		_, err := lambda.
			New(123).
			MapArray(func(idx int, obj interface{}) interface{} { return obj }).
			GroupByArray(func(idx int, obj interface{}) interface{} {
				return obj
			}).
			ToMapInt2Float64Slice()
		as.NotNil(err)
		as.Equal("123(int) can't convert to []interface", err.Error())
	})
}
