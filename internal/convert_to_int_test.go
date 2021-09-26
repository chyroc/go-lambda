package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	as := assert.New(t)

	type args struct {
		v interface{}
	}
	tests := []struct {
		name       string
		args       args
		want       int
		wantErr    bool
		containErr string
	}{
		{
			name: "int-1",
			args: args{v: 1},
			want: 1,
		},
		{
			name:       "str-2",
			args:       args{v: "str-2"},
			wantErr:    true,
			containErr: "str-2(string) can't convert to int",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt(tt.args.v)
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
