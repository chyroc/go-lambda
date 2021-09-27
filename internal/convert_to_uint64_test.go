package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUint64(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       uint64
		wantErr    bool
		containErr string
	}{

		{
			name: "ToUint64 - uint(1)",
			args: uint(1),

			want: 1,
		},

		{
			name: "ToUint64 - uint8(1)",
			args: uint8(1),

			want: 1,
		},

		{
			name: "ToUint64 - uint16(1)",
			args: uint16(1),

			want: 1,
		},

		{
			name: "ToUint64 - uint32(1)",
			args: uint32(1),

			want: 1,
		},

		{
			name: "ToUint64 - uint64(1)",
			args: uint64(1),

			want: 1,
		},

		{
			name: "ToUint64 - str",
			args: "str",

			wantErr:    true,
			containErr: "str(string) can't convert to uint64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUint64(tt.args)
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
