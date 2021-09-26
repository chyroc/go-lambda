package lambda_test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_Err(t *testing.T) {
	as := assert.New(t)
	var err error

	t.Run("", func(t *testing.T) {
		err := lambda.New(1).WithErr(fmt.Errorf("1")).Error()
		as.NotNil(err)
		as.Equal("1", err.Error())
	})

	t.Run("", func(t *testing.T) {
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).FilterArray(nil).Error())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).MapArray(nil).Error())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).ArrayAsync(nil).Error())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).ArrayAsyncWithErr(nil).Error())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).Transfer(nil).Error())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).Flatten().Error())
		_, err = lambda.New(1).WithErr(fmt.Errorf("1")).String()
		as.NotNil(err)
		resp := []int{}
		err = lambda.New(1).WithErr(fmt.Errorf("1")).ToList(&resp)
		as.NotNil(err)
	})

	t.Run("err - not-change-self", func(t *testing.T) {
		req := lambda.New([]int{0, 1, 2})
		req.WithErr(fmt.Errorf("err"))
		res, err := req.ToIntList()
		as.Nil(err)
		as.Equal([]int{0, 1, 2}, res)
	})
}
