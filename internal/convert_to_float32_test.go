package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat32(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       float32
		wantErr    bool
		containErr string
	}{

		{
			name: "ToFloat32 - float32(1)",
			args: float32(1),

			want: 1,
		},

		{
			name: "ToFloat32 - str",
			args: "str",

			wantErr:    true,
			containErr: "str(string) can't convert to float32",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat32(tt.args)
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
