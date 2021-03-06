package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/chyroc/go-lambda/generate-code/internal"
)

// TODO: uintptr

func main() {
	cli := new(GenerateBasicType)

	toTypeFuncs := []string{}
	toTypeFuncTests := []string{}
	for _, req := range BasicToTypeReqs {
		toTypeFunc, err := cli.GenerateToTypeCode(req)
		if err != nil {
			panic(err)
		}
		toTypeFuncs = append(toTypeFuncs, toTypeFunc)

		toTypeFuncTest, err := cli.GenerateToTypeTestCode(req)
		if err != nil {
			panic(err)
		}
		toTypeFuncTests = append(toTypeFuncTests, toTypeFuncTest)
	}

	toTypeFile, err := cli.GenerateToTypeFile(toTypeFuncs)
	if err != nil {
		panic(err)
	}
	toTypeTestFile, err := cli.GenerateToTypeTestFile(toTypeFuncTests)
	if err != nil {
		panic(err)
	}
	toLambdaObjFile, err := cli.GenerateLambdaObjTo(BasicToTypeReqs)
	if err != nil {
		panic(err)
	}
	toLambdaObjTestFile, err := cli.GenerateLambdaObjTestTo(BasicToTypeReqs)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("internal/convert_to_basic.go", []byte(toTypeFile), 0o666); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("internal/convert_to_basic_test.go", []byte(toTypeTestFile), 0o666); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("response_to_basic.go", []byte(toLambdaObjFile), 0o666); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("response_to_basic_test.go", []byte(toLambdaObjTestFile), 0o666); err != nil {
		panic(err)
	}
}

type GenerateBasicType struct{}

func (r *GenerateBasicType) GenerateToTypeCode(req *internal.ToTypeReq) (string, error) {
	tem := `func To{{.TypeTitle}}(v interface{}) ({{.Type}}, error) {
	switch v := v.(type) {
{{range .ConvertTypes}}	case {{.}}:
{{if eq . $.Type}}		return v, nil
{{else}}		return {{$.Type}}(v), nil{{end}}
{{end}}{{range .OverflowTypes}}	case {{.}}:
{{if $.GroupTypeMaxValType}}		if {{$.GroupTypeMaxValType}}(v) <= {{$.MaxVal}}{
			return {{$.Type}}(v), nil
		}
{{else}}		if v <= {{$.MaxVal}}{
			return {{$.Type}}(v), nil
		}
{{end}}	return {{$.ZeroVal}}, fmt.Errorf("%v(%T) overflow max {{$.Type}}", v, v)
{{end}}	default:
		return {{.ZeroVal}}, fmt.Errorf("%v(%T) can't convert to {{.Type}}", v, v)
	}
}
`
	return r.BuildTemplate(tem, req)
}

func (r *GenerateBasicType) GenerateToTypeTestCode(req *internal.ToTypeReq) (string, error) {
	for _, v := range req.TestCases {
		if v.ArgsType == "str" {
			v.ArgsLite = fmt.Sprintf("%q", v.Args)
		} else {
			v.ArgsLite = v.Args
		}
		v.Args = strings.ReplaceAll(v.Args, `"`, "")
	}

	tem := `func TestTo{{.TypeTitle}}(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       {{.Type}}
		errContain string
	}{
{{range .TestCases}}		{name: "To{{$.TypeTitle}} - {{.Args}}", args: {{.ArgsLite}}, {{if .ErrContain}} errContain: "{{.ErrContain}}"{{else}} want: {{.Want}}{{end}}},
{{end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := To{{.TypeTitle}}(tt.args)
			if tt.errContain != "" {
				as.NotNil(err, fmt.Sprintf("%s, got=%v", tt.name, got))
				as.Contains(err.Error(), tt.errContain, fmt.Sprintf("%s, got=%v", tt.name, got))
				return
			}

			as.Nil(err, tt.name)
			as.Equal(tt.want, got, tt.name)
		})
	}
}
`
	return r.BuildTemplate(tem, req)
}

func (r *GenerateBasicType) GenerateToTypeFile(toTypeFunctions []string) (string, error) {
	tem := `package internal

import (
	"fmt"
	"math"
)

{{range .ToTypeFunctions}}{{.}}
{{end}}

`
	data := map[string]interface{}{
		"ToTypeFunctions": toTypeFunctions,
	}
	return r.BuildTemplate(tem, data)
}

func (r *GenerateBasicType) GenerateToTypeTestFile(toTypeFunctionTests []string) (string, error) {
	tem := `package internal

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

{{range .ToTypeFunctionTests}}{{.}}
{{end}}

`
	data := map[string]interface{}{
		"ToTypeFunctionTests": toTypeFunctionTests,
	}
	return r.BuildTemplate(tem, data)
}

func (r *GenerateBasicType) BuildTemplate(tem string, data interface{}) (string, error) {
	t, err := template.New("").Parse(tem)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (r *GenerateBasicType) GenerateLambdaObjTo(toTypeReqs []*internal.ToTypeReq) (string, error) {
	tem := `package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

{{range .ToTypeReqs}}func (r *Object) To{{.TypeTitle}}() ({{.Type}}, error) {
	if r.err != nil {
		return {{.ZeroVal}}, r.err
	}
	return internal.To{{.TypeTitle}}(r.obj)
}
{{end}}
`
	data := map[string]interface{}{
		"ToTypeReqs": toTypeReqs,
	}
	return r.BuildTemplate(tem, data)
}

func (r *GenerateBasicType) GenerateLambdaObjTestTo(toTypeReqs []*internal.ToTypeReq) (string, error) {
	tem := `package lambda_test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_To(t *testing.T) {
	anyVal := item{}
	as := assert.New(t)

{{range	.ToTypeReqs}}t.Run("To{{.TypeTitle}}", func(t *testing.T) {
		t.Run("To{{.TypeTitle}} - success", func(t *testing.T) {
			res, err := lambda.New({{.Type}}({{.OneVal}})).To{{.TypeTitle}}()
			as.Nil(err)
			as.Equal({{.Type}}({{.OneVal}}), res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).To{{.TypeTitle}}()
			as.NotNil(err)
			as.Contains(err.Error(), "can't convert")
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).To{{.TypeTitle}}()
			as.NotNil(err)
		})
	})
{{end}}
}
`
	data := map[string]interface{}{
		"ToTypeReqs": toTypeReqs,
	}
	return r.BuildTemplate(tem, data)
}

var BasicToTypeReqs = []*internal.ToTypeReq{
	// "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"

	// int
	{
		Type:                "int",
		ZeroVal:             "0",
		OneVal:              "int(1)",
		MaxVal:              "math.MaxInt",
		ConvertTypes:        []string{"int", "int8", "int16", "int32", "uint8", "uint16"},
		OverflowTypes:       []string{"uint64", "uint", "int64", "uint32", "uintptr"},
		GroupTypeMaxValType: "uint64",
		TypeTitle:           "Int",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "int(1)"},
			{Args: "int8(1)", Want: "int(1)"},
			{Args: "int16(1)", Want: "int(1)"},
			{Args: "int32(1)", Want: "int(1)"},
			{Args: "int64(1)", Want: "int(1)"},
			{Args: "uint(1)", Want: "int(1)"},
			{Args: "uint8(1)", Want: "int(1)"},
			{Args: "uint16(1)", Want: "int(1)"},
			{Args: "uint32(1)", Want: "int(1)"},
			{Args: "uint64(1)", Want: "int(1)"},
			{Args: "uintptr(1)", Want: "int(1)"},

			// max
			{Args: "int(math.MaxInt)", Want: "int(math.MaxInt)"},
			{Args: "int8(math.MaxInt8)", Want: "int(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", Want: "int(math.MaxInt16)"},
			{Args: "int32(math.MaxInt32)", Want: "int(math.MaxInt32)"},
			{Args: "int64(math.MaxInt64)", Want: "int(math.MaxInt64)"},
			{Args: "uint(math.MaxUint)", ErrContain: "overflow"},
			{Args: "uint8(math.MaxUint8)", Want: "int(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", Want: "int(math.MaxUint16)"},
			{Args: "uint32(math.MaxUint32)", Want: "int(math.MaxUint32)"},
			{Args: "uint64(math.MaxUint64)", ErrContain: "overflow"},
			{Args: "uintptr(math.MaxUint64)", ErrContain: "overflow"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "int8",
		ZeroVal:       "0",
		OneVal:        "int8(1)",
		MaxVal:        "math.MaxInt8",
		ConvertTypes:  []string{"int8"},
		OverflowTypes: []string{"int", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"},
		TypeTitle:     "Int8",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "int8(1)"},
			{Args: "int8(1)", Want: "int8(1)"},
			{Args: "int16(1)", Want: "int8(1)"},
			{Args: "int32(1)", Want: "int8(1)"},
			{Args: "int64(1)", Want: "int8(1)"},
			{Args: "uint(1)", Want: "int8(1)"},
			{Args: "uint8(1)", Want: "int8(1)"},
			{Args: "uint16(1)", Want: "int8(1)"},
			{Args: "uint32(1)", Want: "int8(1)"},
			{Args: "uint64(1)", Want: "int8(1)"},
			{Args: "uintptr(1)", Want: "int8(1)"},

			// max
			{Args: "int(math.MaxInt)", ErrContain: "overflow"},
			{Args: "int8(math.MaxInt8)", Want: "int8(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", ErrContain: "overflow"},
			{Args: "int32(math.MaxInt32)", ErrContain: "overflow"},
			{Args: "int64(math.MaxInt64)", ErrContain: "overflow"},
			{Args: "uint(math.MaxUint)", ErrContain: "overflow"},
			{Args: "uint8(math.MaxUint8)", ErrContain: "overflow"},
			{Args: "uint16(math.MaxUint16)", ErrContain: "overflow"},
			{Args: "uint32(math.MaxUint32)", ErrContain: "overflow"},
			{Args: "uint64(math.MaxUint64)", ErrContain: "overflow"},
			{Args: "uintptr(math.MaxUint64)", ErrContain: "overflow"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "int16",
		ZeroVal:       "0",
		OneVal:        "int16(1)",
		MaxVal:        "math.MaxInt16",
		ConvertTypes:  []string{"int8", "int16", "uint8"},
		OverflowTypes: []string{"int", "int32", "int64", "uint", "uint16", "uint32", "uint64", "uintptr"},
		TypeTitle:     "Int16",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "int16(1)"},
			{Args: "int8(1)", Want: "int16(1)"},
			{Args: "int16(1)", Want: "int16(1)"},
			{Args: "int32(1)", Want: "int16(1)"},
			{Args: "int64(1)", Want: "int16(1)"},
			{Args: "uint(1)", Want: "int16(1)"},
			{Args: "uint8(1)", Want: "int16(1)"},
			{Args: "uint16(1)", Want: "int16(1)"},
			{Args: "uint32(1)", Want: "int16(1)"},
			{Args: "uint64(1)", Want: "int16(1)"},
			{Args: "uintptr(1)", Want: "int16(1)"},

			// max
			{Args: "int(math.MaxInt)", ErrContain: "overflow"},
			{Args: "int8(math.MaxInt8)", Want: "int16(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", Want: "int16(math.MaxInt16)"},
			{Args: "int32(math.MaxInt32)", ErrContain: "overflow"},
			{Args: "int64(math.MaxInt64)", ErrContain: "overflow"},
			{Args: "uint(math.MaxUint)", ErrContain: "overflow"},
			{Args: "uint8(math.MaxUint8)", Want: "int16(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", ErrContain: "overflow"},
			{Args: "uint32(math.MaxUint32)", ErrContain: "overflow"},
			{Args: "uint64(math.MaxUint64)", ErrContain: "overflow"},
			{Args: "uintptr(math.MaxUint64)", ErrContain: "overflow"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "int32",
		ZeroVal:       "0",
		OneVal:        "int32(1)",
		MaxVal:        "math.MaxInt32",
		ConvertTypes:  []string{"int8", "int16", "int32", "uint8", "uint16"},
		OverflowTypes: []string{"int", "int64", "uint", "uint32", "uint64", "uintptr"},
		TypeTitle:     "Int32",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "int32(1)"},
			{Args: "int8(1)", Want: "int32(1)"},
			{Args: "int16(1)", Want: "int32(1)"},
			{Args: "int32(1)", Want: "int32(1)"},
			{Args: "int64(1)", Want: "int32(1)"},
			{Args: "uint(1)", Want: "int32(1)"},
			{Args: "uint8(1)", Want: "int32(1)"},
			{Args: "uint16(1)", Want: "int32(1)"},
			{Args: "uint32(1)", Want: "int32(1)"},
			{Args: "uint64(1)", Want: "int32(1)"},
			{Args: "uintptr(1)", Want: "int32(1)"},

			// max
			{Args: "int(math.MaxInt)", ErrContain: "overflow"},
			{Args: "int8(math.MaxInt8)", Want: "int32(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", Want: "int32(math.MaxInt16)"},
			{Args: "int32(math.MaxInt32)", Want: "int32(math.MaxInt32)"},
			{Args: "int64(math.MaxInt64)", ErrContain: "overflow"},
			{Args: "uint(math.MaxUint)", ErrContain: "overflow"},
			{Args: "uint8(math.MaxUint8)", Want: "int32(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", Want: "int32(math.MaxUint16)"},
			{Args: "uint32(math.MaxUint32)", ErrContain: "overflow"},
			{Args: "uint64(math.MaxUint64)", ErrContain: "overflow"},
			{Args: "uintptr(math.MaxUint64)", ErrContain: "overflow"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "int64",
		ZeroVal:       "0",
		OneVal:        "int64(1)",
		MaxVal:        "math.MaxInt64",
		ConvertTypes:  []string{"int", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32"},
		OverflowTypes: []string{"uint", "uint64", "uintptr"},
		TypeTitle:     "Int64",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "int64(1)"},
			{Args: "int8(1)", Want: "int64(1)"},
			{Args: "int16(1)", Want: "int64(1)"},
			{Args: "int32(1)", Want: "int64(1)"},
			{Args: "int64(1)", Want: "int64(1)"},
			{Args: "uint(1)", Want: "int64(1)"},
			{Args: "uint8(1)", Want: "int64(1)"},
			{Args: "uint16(1)", Want: "int64(1)"},
			{Args: "uint32(1)", Want: "int64(1)"},
			{Args: "uint64(1)", Want: "int64(1)"},
			{Args: "uintptr(1)", Want: "int64(1)"},

			// max
			{Args: "int(math.MaxInt)", Want: "int64(math.MaxInt)"},
			{Args: "int8(math.MaxInt8)", Want: "int64(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", Want: "int64(math.MaxInt16)"},
			{Args: "int32(math.MaxInt32)", Want: "int64(math.MaxInt32)"},
			{Args: "int64(math.MaxInt64)", Want: "int64(math.MaxInt64)"},
			{Args: "uint(math.MaxUint)", ErrContain: "overflow"},
			{Args: "uint8(math.MaxUint8)", Want: "int64(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", Want: "int64(math.MaxUint16)"},
			{Args: "uint32(math.MaxUint32)", Want: "int64(math.MaxUint32)"},
			{Args: "uint64(math.MaxUint64)", ErrContain: "overflow"},
			{Args: "uintptr(math.MaxUint64)", ErrContain: "overflow"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},

	// uint
	{
		Type:                "uint",
		ZeroVal:             "0",
		OneVal:              "uint(1)",
		MaxVal:              "math.MaxUint",
		ConvertTypes:        []string{"int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32"},
		OverflowTypes:       []string{"int64", "uint64", "uintptr"},
		GroupTypeMaxValType: "uint64",
		TypeTitle:           "Uint",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "uint(1)"},
			{Args: "int8(1)", Want: "uint(1)"},
			{Args: "int16(1)", Want: "uint(1)"},
			{Args: "int32(1)", Want: "uint(1)"},
			{Args: "int64(1)", Want: "uint(1)"},
			{Args: "uint(1)", Want: "uint(1)"},
			{Args: "uint8(1)", Want: "uint(1)"},
			{Args: "uint16(1)", Want: "uint(1)"},
			{Args: "uint32(1)", Want: "uint(1)"},
			{Args: "uint64(1)", Want: "uint(1)"},
			{Args: "uintptr(1)", Want: "uint(1)"},

			// max
			{Args: "int(math.MaxInt)", Want: "uint(math.MaxInt)"},
			{Args: "int8(math.MaxInt8)", Want: "uint(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", Want: "uint(math.MaxInt16)"},
			{Args: "int32(math.MaxInt32)", Want: "uint(math.MaxInt32)"},
			// {Args: "int64(math.MaxInt64)", Want: "uint(math.MaxInt64)"},
			{Args: "uint(math.MaxUint)", Want: "uint(math.MaxUint)"},
			{Args: "uint8(math.MaxUint8)", Want: "uint(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", Want: "uint(math.MaxUint16)"},
			{Args: "uint32(math.MaxUint32)", Want: "uint(math.MaxUint32)"},
			{Args: "uint64(math.MaxUint64)", Want: "uint(math.MaxUint64)"},
			{Args: "uintptr(math.MaxUint64)", Want: "uint(math.MaxUint64)"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "uint8",
		ZeroVal:       "0",
		OneVal:        "uint8(1)",
		MaxVal:        "math.MaxUint8",
		ConvertTypes:  []string{"int8", "uint8"},
		OverflowTypes: []string{"int", "int16", "int32", "int64", "uint", "uint16", "uint32", "uint64", "uintptr"},
		TypeTitle:     "Uint8",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "uint8(1)"},
			{Args: "int8(1)", Want: "uint8(1)"},
			{Args: "int16(1)", Want: "uint8(1)"},
			{Args: "int32(1)", Want: "uint8(1)"},
			{Args: "int64(1)", Want: "uint8(1)"},
			{Args: "uint(1)", Want: "uint8(1)"},
			{Args: "uint8(1)", Want: "uint8(1)"},
			{Args: "uint16(1)", Want: "uint8(1)"},
			{Args: "uint32(1)", Want: "uint8(1)"},
			{Args: "uint64(1)", Want: "uint8(1)"},
			{Args: "uintptr(1)", Want: "uint8(1)"},

			// max
			{Args: "int(math.MaxInt)", ErrContain: "overflow"},
			{Args: "int8(math.MaxInt8)", Want: "uint8(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", ErrContain: "overflow"},
			{Args: "int32(math.MaxInt32)", ErrContain: "overflow"},
			{Args: "int64(math.MaxInt64)", ErrContain: "overflow"},
			{Args: "uint(math.MaxUint)", ErrContain: "overflow"},
			{Args: "uint8(math.MaxUint8)", Want: "uint8(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", ErrContain: "overflow"},
			{Args: "uint32(math.MaxUint32)", ErrContain: "overflow"},
			{Args: "uint64(math.MaxUint64)", ErrContain: "overflow"},
			{Args: "uintptr(math.MaxUint64)", ErrContain: "overflow"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "uint16",
		ZeroVal:       "0",
		OneVal:        "uint16(1)",
		MaxVal:        "math.MaxUint16",
		ConvertTypes:  []string{"int8", "int16", "uint8", "uint16"},
		OverflowTypes: []string{"int", "int32", "int64", "uint", "uint32", "uint64", "uintptr"},
		TypeTitle:     "Uint16",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "uint16(1)"},
			{Args: "int8(1)", Want: "uint16(1)"},
			{Args: "int16(1)", Want: "uint16(1)"},
			{Args: "int32(1)", Want: "uint16(1)"},
			{Args: "int64(1)", Want: "uint16(1)"},
			{Args: "uint(1)", Want: "uint16(1)"},
			{Args: "uint8(1)", Want: "uint16(1)"},
			{Args: "uint16(1)", Want: "uint16(1)"},
			{Args: "uint32(1)", Want: "uint16(1)"},
			{Args: "uint64(1)", Want: "uint16(1)"},
			{Args: "uintptr(1)", Want: "uint16(1)"},

			// max
			{Args: "int(math.MaxInt)", ErrContain: "overflow"},
			{Args: "int8(math.MaxInt8)", Want: "uint16(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", Want: "uint16(math.MaxInt16)"},
			{Args: "int32(math.MaxInt32)", ErrContain: "overflow"},
			{Args: "int64(math.MaxInt64)", ErrContain: "overflow"},
			{Args: "uint(math.MaxUint)", ErrContain: "overflow"},
			{Args: "uint8(math.MaxUint8)", Want: "uint16(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", Want: "uint16(math.MaxUint16)"},
			{Args: "uint32(math.MaxUint32)", ErrContain: "overflow"},
			{Args: "uint64(math.MaxUint64)", ErrContain: "overflow"},
			{Args: "uintptr(math.MaxUint64)", ErrContain: "overflow"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "uint32",
		ZeroVal:       "0",
		OneVal:        "uint32(1)",
		MaxVal:        "math.MaxUint32",
		ConvertTypes:  []string{"int8", "int16", "int32", "uint8", "uint16", "uint32"},
		OverflowTypes: []string{"int", "int64", "uint", "uint64", "uintptr"},
		TypeTitle:     "Uint32",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "uint32(1)"},
			{Args: "int8(1)", Want: "uint32(1)"},
			{Args: "int16(1)", Want: "uint32(1)"},
			{Args: "int32(1)", Want: "uint32(1)"},
			{Args: "int64(1)", Want: "uint32(1)"},
			{Args: "uint(1)", Want: "uint32(1)"},
			{Args: "uint8(1)", Want: "uint32(1)"},
			{Args: "uint16(1)", Want: "uint32(1)"},
			{Args: "uint32(1)", Want: "uint32(1)"},
			{Args: "uint64(1)", Want: "uint32(1)"},
			{Args: "uintptr(1)", Want: "uint32(1)"},

			// max
			{Args: "int(math.MaxInt)", ErrContain: "overflow"},
			{Args: "int8(math.MaxInt8)", Want: "uint32(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", Want: "uint32(math.MaxInt16)"},
			{Args: "int32(math.MaxInt32)", Want: "uint32(math.MaxInt32)"},
			{Args: "int64(math.MaxInt64)", ErrContain: "overflow"},
			{Args: "uint(math.MaxUint)", ErrContain: "overflow"},
			{Args: "uint8(math.MaxUint8)", Want: "uint32(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", Want: "uint32(math.MaxUint16)"},
			{Args: "uint32(math.MaxUint32)", Want: "uint32(math.MaxUint32)"},
			{Args: "uint64(math.MaxUint64)", ErrContain: "overflow"},
			{Args: "uintptr(math.MaxUint64)", ErrContain: "overflow"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "uint64",
		ZeroVal:       "0",
		OneVal:        "uint64(1)",
		MaxVal:        "math.MaxUint64",
		ConvertTypes:  []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"},
		OverflowTypes: []string{},
		TypeTitle:     "Uint64",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "uint64(1)"},
			{Args: "int8(1)", Want: "uint64(1)"},
			{Args: "int16(1)", Want: "uint64(1)"},
			{Args: "int32(1)", Want: "uint64(1)"},
			{Args: "int64(1)", Want: "uint64(1)"},
			{Args: "uint(1)", Want: "uint64(1)"},
			{Args: "uint8(1)", Want: "uint64(1)"},
			{Args: "uint16(1)", Want: "uint64(1)"},
			{Args: "uint32(1)", Want: "uint64(1)"},
			{Args: "uint64(1)", Want: "uint64(1)"},
			{Args: "uintptr(1)", Want: "uint64(1)"},

			// max
			{Args: "int(math.MaxInt)", Want: "uint64(math.MaxInt)"},
			{Args: "int8(math.MaxInt8)", Want: "uint64(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", Want: "uint64(math.MaxInt16)"},
			{Args: "int32(math.MaxInt32)", Want: "uint64(math.MaxInt32)"},
			{Args: "int64(math.MaxInt64)", Want: "uint64(math.MaxInt64)"},
			{Args: "uint(math.MaxUint)", Want: "uint64(math.MaxUint)"},
			{Args: "uint8(math.MaxUint8)", Want: "uint64(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", Want: "uint64(math.MaxUint16)"},
			{Args: "uint32(math.MaxUint32)", Want: "uint64(math.MaxUint32)"},
			{Args: "uint64(math.MaxUint64)", Want: "uint64(math.MaxUint64)"},
			{Args: "uintptr(math.MaxUint64)", Want: "uint64(math.MaxUint64)"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "uintptr",
		ZeroVal:       "0",
		OneVal:        "uintptr(1)",
		MaxVal:        "math.MaxUint64",
		ConvertTypes:  []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"},
		OverflowTypes: []string{},
		TypeTitle:     "Uintptr",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "int(1)", Want: "uintptr(1)"},
			{Args: "int8(1)", Want: "uintptr(1)"},
			{Args: "int16(1)", Want: "uintptr(1)"},
			{Args: "int32(1)", Want: "uintptr(1)"},
			{Args: "int64(1)", Want: "uintptr(1)"},
			{Args: "uint(1)", Want: "uintptr(1)"},
			{Args: "uint8(1)", Want: "uintptr(1)"},
			{Args: "uint16(1)", Want: "uintptr(1)"},
			{Args: "uint32(1)", Want: "uintptr(1)"},
			{Args: "uint64(1)", Want: "uintptr(1)"},
			{Args: "uintptr(1)", Want: "uintptr(1)"},

			// max
			{Args: "int(math.MaxInt)", Want: "uintptr(math.MaxInt)"},
			{Args: "int8(math.MaxInt8)", Want: "uintptr(math.MaxInt8)"},
			{Args: "int16(math.MaxInt16)", Want: "uintptr(math.MaxInt16)"},
			{Args: "int32(math.MaxInt32)", Want: "uintptr(math.MaxInt32)"},
			{Args: "int64(math.MaxInt64)", Want: "uintptr(math.MaxInt64)"},
			{Args: "uint(math.MaxUint)", Want: "uintptr(math.MaxUint)"},
			{Args: "uint8(math.MaxUint8)", Want: "uintptr(math.MaxUint8)"},
			{Args: "uint16(math.MaxUint16)", Want: "uintptr(math.MaxUint16)"},
			{Args: "uint32(math.MaxUint32)", Want: "uintptr(math.MaxUint32)"},
			{Args: "uint64(math.MaxUint64)", Want: "uintptr(math.MaxUint64)"},
			{Args: "uintptr(math.MaxUint64)", Want: "uintptr(math.MaxUint64)"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},

	// float32
	{
		Type:          "float32",
		ZeroVal:       "0",
		OneVal:        "float32(1)",
		MaxVal:        "math.MaxFloat32",
		ConvertTypes:  []string{"float32"},
		OverflowTypes: []string{"float64"},
		TypeTitle:     "Float32",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "float32(1)", Want: "float32(1)"},
			{Args: "float64(1)", Want: "float32(1)"},

			// max
			{Args: "float32(math.MaxFloat32)", Want: "float32(math.MaxFloat32)"},
			{Args: "float64(math.MaxFloat64)", ErrContain: "overflow"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:          "float64",
		ZeroVal:       "0",
		OneVal:        "float64(1)",
		MaxVal:        "math.MaxFloat64",
		ConvertTypes:  []string{"float32", "float64"},
		OverflowTypes: []string{},
		TypeTitle:     "Float64",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "float32(1)", Want: "float64(1)"},
			{Args: "float64(1)", Want: "float64(1)"},

			// max
			{Args: "float32(math.MaxFloat32)", Want: "float64(math.MaxFloat32)"},
			{Args: "float64(math.MaxFloat64)", Want: "float64(math.MaxFloat64)"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},

	// bool
	{
		Type:          "bool",
		ZeroVal:       "false",
		OneVal:        "true",
		ConvertTypes:  []string{"bool"},
		OverflowTypes: []string{},
		TypeTitle:     "Bool",
		TestCases: []*internal.TestCase{
			// true
			{Args: "true", Want: "true"},
			{Args: "false", Want: "false"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},

	// complex
	{
		Type:         "complex64",
		ZeroVal:      "0",
		OneVal:       "complex64(1)",
		ConvertTypes: []string{"complex64"},
		TypeTitle:    "Complex64",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "complex64(1)", Want: "complex64(1)"},
			{Args: "complex128(1)", ErrContain: "can't convert"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},
	{
		Type:         "complex128",
		ZeroVal:      "0",
		OneVal:       "complex128(1)",
		ConvertTypes: []string{"complex64", "complex128"},
		TypeTitle:    "Complex128",
		TestCases: []*internal.TestCase{
			// 1
			{Args: "complex64(1)", Want: "complex128(1)"},
			{Args: "complex128(1)", Want: "complex128(1)"},

			// other type
			{Args: "str", ArgsType: "str", ErrContain: "can't convert"},
		},
	},

	{
		Type:         "string",
		ZeroVal:      `""`,
		OneVal:       `"1"`,
		ConvertTypes: []string{"[]rune", "[]byte", "string"},
		TypeTitle:    "String",
		TestCases: []*internal.TestCase{
			// 1
			{Args: `string("1")`, Want: `string("1")`},
			{Args: `[]rune("1")`, Want: `string("1")`},
			{Args: `[]byte("1")`, Want: `string("1")`},

			// other type
			{Args: "1", ErrContain: "can't convert"},
		},
	},
}
