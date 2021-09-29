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
		err := lambda.New(1).WithErr(fmt.Errorf("1")).ToError()
		as.NotNil(err)
		as.Equal("1", err.Error())
	})

	t.Run("", func(t *testing.T) {
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).FilterList(nil).ToError())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).MapList(nil).ToError())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).MapArrayAsync(nil).ToError())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).MapArrayAsyncWithErr(nil).ToError())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).Transfer(nil).ToError())
		as.NotNil(lambda.New(1).WithErr(fmt.Errorf("1")).Flatten().ToError())
		_, err = lambda.New(1).WithErr(fmt.Errorf("1")).ToString()
		as.NotNil(err)
		resp := []int{}
		err = lambda.New(1).WithErr(fmt.Errorf("1")).ToList(&resp)
		as.NotNil(err)
	})

	t.Run("err - not-change-self", func(t *testing.T) {
		req := lambda.New([]int{0, 1, 2})
		req.WithErr(fmt.Errorf("err"))
		res, err := req.ToIntSlice()
		as.Nil(err)
		as.Equal([]int{0, 1, 2}, res)
	})
}
