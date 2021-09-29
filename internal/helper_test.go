package internal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectCanInterface(t *testing.T) {
	as := assert.New(t)

	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "int",
			args: args{v: reflect.ValueOf(1)},
			want: false,
		},
		{
			name: "map",
			args: args{v: reflect.ValueOf(map[string]string{})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as.Equal(tt.want, ReflectCanInterface(tt.args.v))
		})
	}
}
