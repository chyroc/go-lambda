package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInterfaceList(t *testing.T) {
	as := assert.New(t)

	type args struct {
		v interface{}
	}
	tests := []struct {
		name       string
		args       args
		want       []interface{}
		wantErr    bool
		containErr string
	}{
		{
			name: "int-list",
			args: args{v: []int{1, 2}},
			want: []interface{}{1, 2},
		},
		{
			name: "int-array",
			args: args{v: [2]int{1, 2}},
			want: []interface{}{1, 2},
		},
		{
			name: "interface-int-list",
			args: args{v: []interface{}{1, 2}},
			want: []interface{}{1, 2},
		},
		{
			name: "str-list",
			args: args{v: []interface{}{"str"}},
			want: []interface{}{"str"},
		},
		{
			name: "str-array",
			args: args{v: [2]interface{}{"str1", "str2"}},
			want: []interface{}{"str1", "str2"},
		},
		{
			name: "str",
			args: args{v: "str"},
			want: []interface{}{'s', 't', 'r'},
		},
		{
			name:       "int",
			args:       args{v: 1},
			wantErr:    true,
			containErr: "1(int) can't convert to []interface",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInterfaceList(tt.args.v)
			if tt.wantErr {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.containErr, tt.name)
				return
			}
			as.Equal(tt.want, got, tt.name)
		})
	}
}
