package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt32(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       int32
		wantErr    bool
		containErr string
	}{

		{
			name: "ToInt32 - int(1)",
			args: int(1),

			want: 1,
		},

		{
			name: "ToInt32 - int8(1)",
			args: int8(1),

			want: 1,
		},

		{
			name: "ToInt32 - int16(1)",
			args: int16(1),

			want: 1,
		},

		{
			name: "ToInt32 - int32(1)",
			args: int32(1),

			want: 1,
		},

		{
			name: "ToInt32 - str",
			args: "str",

			wantErr:    true,
			containErr: "str(string) can't convert to int32",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt32(tt.args)
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
