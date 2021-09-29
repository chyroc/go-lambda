package lambda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_obj(t *testing.T) {
	as := assert.New(t)
	as.Equal(objType(0), New(1).objType())
	as.Equal(objTypeArray, New([]int{1, 2, 3}).objType())
	as.Equal(objTypeMap, New(map[string]string{"1": "2"}).objType())
}
