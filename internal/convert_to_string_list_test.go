package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	as := assert.New(t)

	type args struct {
		v interface{}
	}
	tests := []struct {
		name       string
		args       args
		want       string
		wantErr    bool
		containErr string
	}{
		{
			name: "str",
			args: args{v: "1"},
			want: "1",
		},
		{
			name: "[]byte",
			args: args{v: []byte("1")},
			want: "1",
		},
		{
			name: "[]rune",
			args: args{v: []rune("1")},
			want: "1",
		},
		{
			name:       "[]int64",
			args:       args{v: []int64{1}},
			wantErr:    true,
			containErr: "[1]([]int64) can't convert string",
		},
		{
			name:       "int",
			args:       args{v: 1},
			wantErr:    true,
			containErr: "1(int) can't convert string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.args.v)
			if tt.wantErr {
				as.NotNil(err)
				as.Contains(err.Error(), tt.containErr)
				return
			}

			as.Nil(err)
			as.Equal(tt.want, got)
		})
	}
}
