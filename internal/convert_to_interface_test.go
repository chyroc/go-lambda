package internal

import (
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/stretchr/testify/assert"
)

type toExpectTypeInterfaceItem struct {
	Name        string
	Age         int
	CompanyList []*Company
}

type Company struct {
	Title string
}

func Test_ToInterface(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		etype      interface{}
		want       interface{}
		containErr string
	}{
		// basic
		{name: "bool", args: bool(true), etype: bool(false), want: bool(true)},
		{name: "int", args: int(1), etype: int(0), want: int(1)},
		{name: "int8", args: int8(1), etype: int8(0), want: int8(1)},
		{name: "int16", args: int16(1), etype: int16(0), want: int16(1)},
		{name: "int32", args: int32(1), etype: int32(0), want: int32(1)},
		{name: "int64", args: int64(1), etype: int64(0), want: int64(1)},
		{name: "uint", args: uint(1), etype: uint(0), want: uint(1)},
		{name: "uint8", args: uint8(1), etype: uint8(0), want: uint8(1)},
		{name: "uint16", args: uint16(1), etype: uint16(0), want: uint16(1)},
		{name: "uint32", args: uint32(1), etype: uint32(0), want: uint32(1)},
		{name: "uint64", args: uint64(1), etype: uint64(0), want: uint64(1)},
		{name: "uintptr", args: uintptr(1), etype: uintptr(0), want: uintptr(1)},
		{name: "float32", args: float32(1), etype: float32(0), want: float32(1)},
		{name: "float64", args: float64(1), etype: float64(0), want: float64(1)},
		{name: "complex64", args: complex64(1), etype: complex64(0), want: complex64(1)},
		{name: "complex128", args: complex128(1), etype: complex128(0), want: complex128(1)},
		{name: "string", args: "str", etype: "", want: "str"},

		// to-slice
		{name: "slice-int", args: []int{1}, etype: []int{}, want: []int{1}},
		{name: "slice-int", args: [2]int{1, 2}, etype: []int{}, want: []int{1, 2}},
		{name: "slice-int", args: [2]interface{}{1, 2}, etype: []int{}, want: []int{1, 2}},
		{name: "slice-int", args: [2]interface{}{1, "str"}, etype: []int{}, containErr: "str(string) can't convert to int"},

		// to-array
		{name: "slice-int", args: []int{1}, etype: [1]int{}, want: [1]int{1}},
		{name: "slice-int", args: [2]int{1, 2}, etype: [1]int{}, want: [2]int{1, 2}},
		{name: "slice-int", args: [2]interface{}{1, 2}, etype: [1]int{}, want: [2]int{1, 2}},
		{name: "slice-int", args: [2]interface{}{1, "str"}, etype: [1]int{}, containErr: "str(string) can't convert to int"},

		// map
		{name: "map-1-1", args: map[int]int{1: 1}, etype: map[int]int{}, want: map[int]int{1: 1}},
		{name: "map-1-1", args: map[int]int{1: 1}, etype: map[int]int8{}, want: map[int]int8{1: 1}},
		{name: "map-1-1", args: map[int]int{1: 1}, etype: map[int32]int8{}, want: map[int32]int8{1: 1}},
		{name: "map-1-1", args: map[int][]int{1: {1, 2}}, etype: map[int][2]int{}, want: map[int][2]int{1: {1, 2}}},

		{name: "ptr-1", args: 1, etype: ptr.Int(0), want: ptr.Int(1)},

		// struct
		{
			name: "struct-1", args: toExpectTypeInterfaceItem{Name: "name", Age: 12, CompanyList: []*Company{{Title: "t1"}, {Title: "t2"}}},
			etype: toExpectTypeInterfaceItem{},
			want:  toExpectTypeInterfaceItem{Name: "name", Age: 12, CompanyList: []*Company{{Title: "t1"}, {Title: "t2"}}},
		},
		{
			name: "struct-1", args: &toExpectTypeInterfaceItem{Name: "name", Age: 12, CompanyList: []*Company{{Title: "t1"}, {Title: "t2"}}},
			etype: toExpectTypeInterfaceItem{},
			want:  toExpectTypeInterfaceItem{Name: "name", Age: 12, CompanyList: []*Company{{Title: "t1"}, {Title: "t2"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToExpectTypeInterface(tt.args, tt.etype)
			if tt.containErr != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.containErr, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}
