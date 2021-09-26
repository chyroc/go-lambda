package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat(t *testing.T) {
	as := assert.New(t)

	type args struct {
		v interface{}
	}
	tests := []struct {
		name       string
		args       args
		want       float64
		wantErr    bool
		containErr string
	}{
		{
			name: "float32",
			args: args{v: float32(1)},
			want: 1,
		},
		{
			name: "float64",
			args: args{v: float64(1)},
			want: 1,
		},
		{
			name:       "err",
			args:       args{v: "str"},
			wantErr:    true,
			containErr: "str(string) can't convert to float64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat(tt.args.v)
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
