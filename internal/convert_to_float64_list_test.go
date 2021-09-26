package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat64List(t *testing.T) {
	as := assert.New(t)

	type args struct {
		v interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResp   []float64
		wantErr    bool
		containErr string
	}{
		{
			name:     "float64-list",
			args:     args{v: []float64{1, 2}},
			wantResp: []float64{1, 2},
		},
		{
			name:     "interface-float64-list",
			args:     args{v: []interface{}{float64(1), float64(2)}},
			wantResp: []float64{1, 2},
		},
		{
			name:     "interface-float32-list",
			args:     args{v: []interface{}{float32(1), float32(2)}},
			wantResp: []float64{1, 2},
		},
		{
			name:       "interface-str-list",
			args:       args{v: []interface{}{float32(1), "str"}},
			wantErr:    true,
			containErr: "str(string) can't convert to float64",
		},
		{
			name:       "str",
			args:       args{v: "str"},
			wantErr:    true,
			containErr: "str(string) can't convert to []float64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := ToFloat64List(tt.args.v)
			if tt.wantErr {
				as.NotNil(err)
				as.Contains(err.Error(), tt.containErr)
				return
			}

			as.Nil(err)
			as.Equal(tt.wantResp, gotResp)
		})
	}
}
