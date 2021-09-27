package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUint(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       uint
		wantErr    bool
		containErr string
	}{

		{
			name: "ToUint - uint(1)",
			args: uint(1),

			want: 1,
		},

		{
			name: "ToUint - str",
			args: "str",

			wantErr:    true,
			containErr: "str(string) can't convert to uint",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint(tt.args)
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
