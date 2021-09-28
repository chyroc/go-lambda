package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToIntList(t *testing.T) {
	as := assert.New(t)

	type args struct {
		v interface{}
	}
	tests := []struct {
		name       string
		args       args
		want       []int
		wantErr    bool
		containErr string
	}{
		{
			name: "int-list",
			args: args{v: []int{1, 2}},
			want: []int{1, 2},
		},
		{
			name: "interface-int-list",
			args: args{v: []interface{}{1, 2}},
			want: []int{1, 2},
		},
		{
			name:       "err",
			args:       args{v: []interface{}{"str"}},
			wantErr:    true,
			containErr: "str(string) can't convert to int",
		},
		{
			name:       "err",
			args:       args{v: []string{"str"}},
			wantErr:    true,
			containErr: "[str]([]string) can't convert to []int",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToIntSlice(tt.args.v)
			if tt.wantErr {
				as.NotNil(err)
				as.Contains(err.Error(), tt.containErr)
				return
			}
			as.Equal(tt.want, got)
		})
	}
}
