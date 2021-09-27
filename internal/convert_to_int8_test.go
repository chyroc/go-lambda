package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt8(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       int8
		wantErr    bool
		containErr string
	}{

		{
			name: "ToInt8 - int(1)",
			args: int(1),

			want: 1,
		},

		{
			name: "ToInt8 - int8(1)",
			args: int8(1),

			want: 1,
		},

		{
			name: "ToInt8 - str",
			args: "str",

			wantErr:    true,
			containErr: "str(string) can't convert to int8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt8(tt.args)
			if tt.wantErr {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.containErr, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}
