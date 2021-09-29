package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/chyroc/go-lambda/generate-code/internal"
)

// TODO: uintptr

func main() {
	cli := new(GenerateBasicTypeSlice)

	toTypeSliceFile, err := cli.GenerateToTypeSlice(BasicToTypeListReqs)
	if err != nil {
		panic(err)
	}
	toTypeSliceTestFile, err := cli.GenerateToTypeSliceTest(BasicToTypeListReqs)
	if err != nil {
		panic(err)
	}

	toTypeSliceObjectFile, err := cli.GenerateToTypeSliceObject(BasicToTypeListReqs)
	if err != nil {
		panic(err)
	}
	toTypeSliceObjectTestFile, err := cli.GenerateToTypeSliceObjectTest(BasicToTypeListReqs)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("internal/convert_to_basic_slice.go", []byte(toTypeSliceFile), 0o666); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("internal/convert_to_basic_slice_test.go", []byte(toTypeSliceTestFile), 0o666); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("response_to_basic_list.go", []byte(toTypeSliceObjectFile), 0o666); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("response_to_basic_list_test.go", []byte(toTypeSliceObjectTestFile), 0o666); err != nil {
		panic(err)
	}
}

type GenerateBasicTypeSlice struct{}

func (r *GenerateBasicTypeSlice) GenerateToTypeSlice(toTypeReqs []*internal.ToTypeReq) (string, error) {
	tem := `package internal

import (
	"fmt"
	"reflect"
)

{{range .ToTypeReqs}}func To{{.TypeTitle}}Slice(v interface{}) (resp []{{.Type}}, err error) {
	switch v := v.(type) {
	case []{{.Type}}:
		return v, nil
	case []interface{}:
		for _, vv := range v {
			vvv, err := To{{.TypeTitle}}(vv)
			if err != nil {
				return nil, err
			}
			resp = append(resp, vvv)
		}
		return resp, nil
	default:
		vv := reflect.ValueOf(v)
		catchErr := false
		if vv.Kind() == reflect.Array {
			for i := 0; i < vv.Len(); i++ {
				ii, err := To{{.TypeTitle}}(vv.Index(i).Interface())
				if err != nil {
					catchErr = true
					break
				}
				resp = append(resp, ii)
			}
			if !catchErr {
				return resp, nil
			}
		}
		return nil, fmt.Errorf("%v(%T) can't convert to []{{.Type}}", v, v)
	}
}
{{end}}

`
	data := map[string]interface{}{
		"ToTypeReqs": toTypeReqs,
	}
	return internal.BuildTemplate(tem, data)
}

func (r *GenerateBasicTypeSlice) GenerateToTypeSliceTest(toTypeReqs []*internal.ToTypeReq) (string, error) {
	for _, req := range toTypeReqs {
		for _, v := range req.TestCases {
			if v.ArgsType == "str" {
				v.ArgsLite = fmt.Sprintf("%q", v.Args)
			} else {
				v.ArgsLite = v.Args
			}
			v.Args = strings.ReplaceAll(v.Args, `"`, "-")
			if v.Want != "" {
				if v.ArgsType == "str" {
					v.WantLite = fmt.Sprintf("%q", v.Want)
				} else {
					v.WantLite = v.Want
				}
			}

		}
	}

	tem := `package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

{{range $idx, $val := .ToTypeReqs}}func Test_To{{.TypeTitle}}Slice(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       []{{.Type}}
		errContain string
	}{
{{range .TestCases}}		{name: "To{{$val.TypeTitle}}Slice - {{.Args}}", args: {{.ArgsLite}}, {{if .ErrContain}} errContain: "{{.ErrContain}}"{{else}} want: {{.Want}}{{end}}},
{{end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := To{{.TypeTitle}}Slice(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Contains(err.Error(), tt.errContain, tt.name)
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}
{{end}}
`

	return internal.BuildTemplate(tem, map[string]interface{}{
		"ToTypeReqs": toTypeReqs,
	})
}

func (r *GenerateBasicTypeSlice) GenerateToTypeSliceObject(toTypeReqs []*internal.ToTypeReq) (string, error) {
	tem := `package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

{{range .ToTypeReqs}}func (r *Object) To{{.TypeTitle}}Slice() ([]{{.Type}}, error) {
	if r.err != nil {
		return nil, r.err
	}
	return internal.To{{.TypeTitle}}Slice(r.obj)
}
{{end}}
`
	data := map[string]interface{}{
		"ToTypeReqs": toTypeReqs,
	}
	return internal.BuildTemplate(tem, data)
}

func (r *GenerateBasicTypeSlice) GenerateToTypeSliceObjectTest(toTypeReqs []*internal.ToTypeReq) (string, error) {
	tem := `package lambda_test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

type itemSlice struct{}

func Test_ToBasicSlice(t *testing.T) {
	anyVal := itemSlice{}
	as := assert.New(t)

{{range	$idx, $val := .ToTypeReqs}}t.Run("To{{$val.TypeTitle}}Slice", func(t *testing.T) {
{{range .TestCases}}{{ if .Want}}
		t.Run("success - {{.Args}}", func(t *testing.T) {
			res, err := lambda.New({{.ArgsLite}}).To{{$val.TypeTitle}}Slice()
			as.Nil(err)
			as.Equal({{.Want}}, res)
		})
{{end}}
{{end}}
		t.Run("fail-1", func(t *testing.T) {
			_, err := lambda.New(anyVal).To{{$val.TypeTitle}}Slice()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("fail-2", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).To{{$val.TypeTitle}}Slice()
			as.NotNil(err)
		})
	})
{{end}}
}
`
	data := map[string]interface{}{
		"ToTypeReqs": toTypeReqs,
	}
	return internal.BuildTemplate(tem, data)
}

var BasicToTypeListReqs = []*internal.ToTypeReq{
	// "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"

	// int
	{
		Type:      "int",
		ZeroVal:   "0",
		OneVal:    "int(1)",
		TypeTitle: "Int",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]int{int(1)}", Want: "[]int{int(1)}"},
			{Args: "[]interface{}{int(1)}", Want: "[]int{int(1)}"},
			{Args: "[2]int{int(1), int(2)}", Want: "[]int{int(1), int(2)}"},
			{Args: "[2]interface{}{int(1), int(2)}", Want: "[]int{int(1), int(2)}"},

			// other type
			{Args: "str", ArgsType: `str`, ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "int8",
		ZeroVal:   "0",
		OneVal:    "int8(1)",
		TypeTitle: "Int8",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]int8{int8(1)}", Want: "[]int8{int8(1)}"},
			{Args: "[]interface{}{int8(1)}", Want: "[]int8{int8(1)}"},
			{Args: "[2]int8{int8(1), int8(2)}", Want: "[]int8{int8(1), int8(2)}"},
			{Args: "[2]interface{}{int8(1), int8(2)}", Want: "[]int8{int8(1), int8(2)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "int16",
		ZeroVal:   "0",
		OneVal:    "int16(1)",
		TypeTitle: "Int16",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]int16{int16(1)}", Want: "[]int16{int16(1)}"},
			{Args: "[]interface{}{int16(1)}", Want: "[]int16{int16(1)}"},
			{Args: "[2]int16{int16(1), int16(1)}", Want: "[]int16{int16(1), int16(1)}"},
			{Args: "[2]interface{}{int16(1), int16(1)}", Want: "[]int16{int16(1), int16(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "int32",
		ZeroVal:   "0",
		OneVal:    "int32(1)",
		TypeTitle: "Int32",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]int32{int32(1)}", Want: "[]int32{int32(1)}"},
			{Args: "[]interface{}{int32(1)}", Want: "[]int32{int32(1)}"},
			{Args: "[2]int32{int32(1), int32(1)}", Want: "[]int32{int32(1), int32(1)}"},
			{Args: "[2]interface{}{int32(1), int32(1)}", Want: "[]int32{int32(1), int32(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "int64",
		ZeroVal:   "0",
		OneVal:    "int64(1)",
		TypeTitle: "Int64",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]int64{int64(1)}", Want: "[]int64{int64(1)}"},
			{Args: "[]interface{}{int64(1)}", Want: "[]int64{int64(1)}"},
			{Args: "[2]int64{int64(1), int64(1)}", Want: "[]int64{int64(1), int64(1)}"},
			{Args: "[2]interface{}{int64(1), int64(1)}", Want: "[]int64{int64(1), int64(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "uint",
		ZeroVal:   "0",
		OneVal:    "uint(1)",
		TypeTitle: "Uint",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]uint{uint(1)}", Want: "[]uint{uint(1)}"},
			{Args: "[]interface{}{uint(1)}", Want: "[]uint{uint(1)}"},
			{Args: "[2]uint{uint(1), uint(1)}", Want: "[]uint{uint(1), uint(1)}"},
			{Args: "[2]interface{}{uint(1), uint(1)}", Want: "[]uint{uint(1), uint(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "uint8",
		ZeroVal:   "0",
		OneVal:    "uint8(1)",
		TypeTitle: "Uint8",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]uint8{uint8(1)}", Want: "[]uint8{uint8(1)}"},
			{Args: "[]interface{}{uint8(1)}", Want: "[]uint8{uint8(1)}"},
			{Args: "[2]uint8{uint8(1), uint8(1)}", Want: "[]uint8{uint8(1), uint8(1)}"},
			{Args: "[2]interface{}{uint8(1), uint8(1)}", Want: "[]uint8{uint8(1), uint8(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "uint16",
		ZeroVal:   "0",
		OneVal:    "uint16(1)",
		TypeTitle: "Uint16",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]uint16{uint16(1)}", Want: "[]uint16{uint16(1)}"},
			{Args: "[]interface{}{uint16(1)}", Want: "[]uint16{uint16(1)}"},
			{Args: "[2]uint16{uint16(1), uint16(1)}", Want: "[]uint16{uint16(1), uint16(1)}"},
			{Args: "[2]interface{}{uint16(1), uint16(1)}", Want: "[]uint16{uint16(1), uint16(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "uint32",
		ZeroVal:   "0",
		OneVal:    "uint32(1)",
		TypeTitle: "Uint32",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]uint32{uint32(1)}", Want: "[]uint32{uint32(1)}"},
			{Args: "[]interface{}{uint32(1)}", Want: "[]uint32{uint32(1)}"},
			{Args: "[2]uint32{uint32(1), uint32(1)}", Want: "[]uint32{uint32(1), uint32(1)}"},
			{Args: "[2]interface{}{uint32(1), uint32(1)}", Want: "[]uint32{uint32(1), uint32(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "uint64",
		ZeroVal:   "0",
		OneVal:    "uint64(1)",
		TypeTitle: "Uint64",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]uint64{uint64(1)}", Want: "[]uint64{uint64(1)}"},
			{Args: "[]interface{}{uint64(1)}", Want: "[]uint64{uint64(1)}"},
			{Args: "[2]uint64{uint64(1), uint64(1)}", Want: "[]uint64{uint64(1), uint64(1)}"},
			{Args: "[2]interface{}{uint64(1), uint64(1)}", Want: "[]uint64{uint64(1), uint64(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},

	// float32
	{
		Type:      "float32",
		ZeroVal:   "0",
		OneVal:    "float32(1)",
		TypeTitle: "Float32",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]float32{float32(1)}", Want: "[]float32{float32(1)}"},
			{Args: "[]interface{}{float32(1)}", Want: "[]float32{float32(1)}"},
			{Args: "[2]float32{float32(1), float32(1)}", Want: "[]float32{float32(1), float32(1)}"},
			{Args: "[2]interface{}{float32(1), float32(1)}", Want: "[]float32{float32(1), float32(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "float64",
		ZeroVal:   "0",
		OneVal:    "float64(1)",
		TypeTitle: "Float64",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]float64{float64(1)}", Want: "[]float64{float64(1)}"},
			{Args: "[]interface{}{float64(1)}", Want: "[]float64{float64(1)}"},
			{Args: "[2]float64{float64(1), float64(1)}", Want: "[]float64{float64(1), float64(1)}"},
			{Args: "[2]interface{}{float64(1), float64(1)}", Want: "[]float64{float64(1), float64(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},

	// bool
	{
		Type:      "bool",
		ZeroVal:   "false",
		OneVal:    "bool(true)",
		TypeTitle: "Bool",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]bool{bool(true)}", Want: "[]bool{bool(true)}"},
			{Args: "[]interface{}{bool(true)}", Want: "[]bool{bool(true)}"},
			{Args: "[2]bool{bool(true), bool(true)}", Want: "[]bool{bool(true), bool(true)}"},
			{Args: "[2]interface{}{bool(true), bool(true)}", Want: "[]bool{bool(true), bool(true)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},

	// complex
	{
		Type:      "complex64",
		ZeroVal:   "0",
		OneVal:    "complex64(1)",
		TypeTitle: "Complex64",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]complex64{complex64(1)}", Want: "[]complex64{complex64(1)}"},
			{Args: "[]interface{}{complex64(1)}", Want: "[]complex64{complex64(1)}"},
			{Args: "[2]complex64{complex64(1), complex64(1)}", Want: "[]complex64{complex64(1), complex64(1)}"},
			{Args: "[2]interface{}{complex64(1), complex64(1)}", Want: "[]complex64{complex64(1), complex64(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},
	{
		Type:      "complex128",
		ZeroVal:   "0",
		OneVal:    "complex128(1)",
		TypeTitle: "Complex128",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "[]complex128{complex128(1)}", Want: "[]complex128{complex128(1)}"},
			{Args: "[]interface{}{complex128(1)}", Want: "[]complex128{complex128(1)}"},
			{Args: "[2]complex128{complex128(1), complex128(1)}", Want: "[]complex128{complex128(1), complex128(1)}"},
			{Args: "[2]interface{}{complex128(1), complex128(1)}", Want: "[]complex128{complex128(1), complex128(1)}"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]string{"str"}`, ErrContain: "can't convert"},
		},
	},

	{
		Type:      "string",
		ZeroVal:   `""`,
		OneVal:    `"1"`,
		TypeTitle: "String",
		TestCases: []*internal.TestCase{
			// 1
			{Args: `[]string{"1"}`, Want: `[]string{"1"}`},
			{Args: `[]interface{}{"1"}`, Want: `[]string{"1"}`},
			{Args: `[2]string{"1", "2"}`, Want: `[]string{"1", "2"}`},
			{Args: `[2]interface{}{"1", "2"}`, Want: `[]string{"1", "2"}`},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
			{Args: `[]int{1}`, ErrContain: "can't convert"},
		},
	},
}
