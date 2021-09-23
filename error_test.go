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
		err := lambda.New(1).SetErr(fmt.Errorf("1")).Error()
		as.NotNil(err)
		as.Equal("1", err.Error())
	})

	t.Run("", func(t *testing.T) {
		as.NotNil(lambda.New(1).SetErr(fmt.Errorf("1")).Filter(nil).Error())
		as.NotNil(lambda.New(1).SetErr(fmt.Errorf("1")).Map(nil).Error())
		as.NotNil(lambda.New(1).SetErr(fmt.Errorf("1")).Transfer(nil).Error())
		_, err = lambda.New(1).SetErr(fmt.Errorf("1")).String()
		as.NotNil(err)
		resp := []int{}
		err = lambda.New(1).SetErr(fmt.Errorf("1")).ToList(&resp)
		as.NotNil(err)
	})
}
