package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
	as := assert.New(t)

	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// err
		{name: "{v: fmt.Errorf(\"err\")}", args: args{v: fmt.Errorf("err")}, want: false},
		{name: "{v: (error)(nil)}", args: args{v: (error)(nil)}, want: true},

		// int
		{name: "{v: int(1)}", args: args{v: int(1)}, want: false},
		{name: "{v: int8(1)}", args: args{v: int8(1)}, want: false},
		{name: "{v: int16(1)}", args: args{v: int16(1)}, want: false},
		{name: "{v: int32(1)}", args: args{v: int32(1)}, want: false},
		{name: "{v: int64(1)}", args: args{v: int64(1)}, want: false},
		{name: "{v: uint(1)}", args: args{v: uint(1)}, want: false},
		{name: "{v: uint8(1)}", args: args{v: uint8(1)}, want: false},
		{name: "{v: uint16(1)}", args: args{v: uint16(1)}, want: false},
		{name: "{v: uint32(1)}", args: args{v: uint32(1)}, want: false},
		{name: "{v: uint64(1)}", args: args{v: uint64(1)}, want: false},
		{name: "{v: uintptr(1)}", args: args{v: uintptr(1)}, want: false},

		{name: "{v: int(0)}", args: args{v: int(0)}, want: true},
		{name: "{v: int8(0)}", args: args{v: int8(0)}, want: true},
		{name: "{v: int16(0)}", args: args{v: int16(0)}, want: true},
		{name: "{v: int32(0)}", args: args{v: int32(0)}, want: true},
		{name: "{v: int64(0)}", args: args{v: int64(0)}, want: true},
		{name: "{v: uint(0)}", args: args{v: uint(0)}, want: true},
		{name: "{v: uint8(0)}", args: args{v: uint8(0)}, want: true},
		{name: "{v: uint16(0)}", args: args{v: uint16(0)}, want: true},
		{name: "{v: uint32(0)}", args: args{v: uint32(0)}, want: true},
		{name: "{v: uint64(0)}", args: args{v: uint64(0)}, want: true},
		{name: "{v: uintptr(0)}", args: args{v: uintptr(0)}, want: true},

		// float64(),
		{name: "{v: float32(1)}", args: args{v: float32(1)}, want: false},
		{name: "{v: float64(1)}", args: args{v: float64(1)}, want: false},

		{name: "{v: float32(0)}", args: args{v: float32(0)}, want: true},
		{name: "{v: float64(0)}", args: args{v: float64(0)}, want: true},

		// str
		{name: "{v: \"str\"}", args: args{v: "str"}, want: false},
		{name: "{v: \"\"}", args: args{v: ""}, want: true},

		// bool
		{name: "{v: true}", args: args{v: true}, want: false},
		{name: "{v: false}", args: args{v: false}, want: true},

		// interface
		{name: "{v: interface{}(1)}", args: args{v: interface{}(1)}, want: false},
		{name: "{v: nilInterface}", args: args{v: nilInterface}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as.Equal(tt.want, IsZero(tt.args.v), tt.name)
		})
	}
}

var nilInterface interface{} = nil
