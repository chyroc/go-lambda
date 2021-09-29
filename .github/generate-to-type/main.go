package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"
)

func main() {
	res := []generateToTypeReq{}

	res = append(res, convertCaseList([]string{"int", "int8", "int16", "int32", "int64"}, "0")...)
	res = append(res, convertCaseList([]string{"uint", "uint8", "uint16", "uint32", "uint64"}, "0")...)
	res = append(res, convertCaseList([]string{"float32", "float64"}, "0")...)

	writegenerateToTypeObjTo(res)
	writegenerateToTypeObjToTest(res)
}

func convertCaseList(intTypes []string, defaultValue string) (res []generateToTypeReq) {
	for i := 0; i < len(intTypes); i++ {
		typ := intTypes[i]
		caseTypes := intTypes[:i+1]
		testCases := []testCase{}
		for _, caseType := range caseTypes {
			testCases = append(testCases, testCase{Input: caseType + "(1)", Output: "1"})
		}
		testCases = append(testCases, testCase{Input: "str", InputType: "str", ContainErr: "str(string) can't convert to " + typ})
		req := generateToTypeReq{
			Type:      typ,
			Default:   defaultValue,
			CaseTypes: caseTypes,
			TestCase:  testCases,
		}

		text, err := generateToType(typ, req)
		assert(err)
		assert(ioutil.WriteFile(fmt.Sprintf("internal/convert_to_basic_%s.go", typ), []byte(text), 0o666))

		textTest, err := generateToTypeTest(typ, req)
		assert(err)
		assert(ioutil.WriteFile(fmt.Sprintf("internal/convert_to_basic_%s_test.go", typ), []byte(textTest), 0o666))

		res = append(res, req)
	}

	return res
}

type generateToTypeReq struct {
	Type      string
	Default   string
	CaseTypes []string
	TestCase  []testCase
}

type testCase struct {
	Name       string
	Input      string
	InputType  string
	Output     string
	ContainErr string
}

func writegenerateToTypeObjTo(res []generateToTypeReq) {
	buf := strings.Builder{}
	buf.WriteString(`package lambda

import (
	"github.com/chyroc/go-lambda/internal"
)

`)
	for _, v := range res {
		text, err := generateToTypeObjTo(v.Type, v)
		assert(err)

		buf.WriteString(text)
	}
	assert(ioutil.WriteFile("response_to.go", []byte(buf.String()), 0o666))
}

func writegenerateToTypeObjToTest(res []generateToTypeReq) {
	buf := strings.Builder{}
	buf.WriteString(`package lambda_test

import (
	"testing"
	"fmt"

	"github.com/chyroc/go-lambda"
	"github.com/stretchr/testify/assert"
)

func Test_To(t *testing.T) {
	var anyVal = item{}
	as := assert.New(t)

`)
	for _, v := range res {
		text, err := generateToTypeObjToTest(v.Type, v)
		assert(err)

		buf.WriteString(text)
	}
	buf.WriteString("}")
	assert(ioutil.WriteFile("response_to_test.go", []byte(buf.String()), 0o666))
}

func generateToTypeObjToTest(typ string, req generateToTypeReq) (string, error) {
	TypeTitle := strings.Title(typ)
	return buildTem(temObjToTest, map[string]interface{}{
		"TypeTitle":       TypeTitle,
		"Type":            typ,
		"CaseTypes":       req.CaseTypes,
		"SuccessTestCase": req.TestCase[len(req.TestCase)-2],
		"Default":         req.Default,
	}), nil
}

func generateToTypeObjTo(typ string, req generateToTypeReq) (string, error) {
	TypeTitle := strings.Title(typ)
	return buildTem(temObjTo, map[string]interface{}{
		"TypeTitle": TypeTitle,
		"Type":      typ,
		"CaseTypes": req.CaseTypes,
		"Default":   req.Default,
	}), nil
}

func generateToType(typ string, req generateToTypeReq) (string, error) {
	TypeTitle := strings.Title(typ)
	return buildTem(tem, map[string]interface{}{
		"TypeTitle": TypeTitle,
		"Type":      typ,
		"CaseTypes": req.CaseTypes,
		"Default":   req.Default,
	}), nil
}

func generateToTypeTest(typ string, req generateToTypeReq) (string, error) {
	testCase := []testCase{}
	for _, v := range req.TestCase {
		v.Name = fmt.Sprintf("To%s - %v", strings.Title(typ), v.Input)
		testCase = append(testCase, v)
	}
	return buildTem(temTest, map[string]interface{}{
		"TypeTitle": strings.Title(typ),
		"Type":      typ,
		"CaseTypes": req.CaseTypes,
		"Default":   req.Default,
		"TestCase":  testCase,
	}), nil
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

var tem = `package internal

import (
	"fmt"
)

func To{{.TypeTitle}}(v interface{}) ({{.Type}}, error) {
	switch v := v.(type) {
{{range .CaseTypes}}case {{.}}:
		return {{if eq $.Type .}}v{{else}}{{$.Type}}(v){{end}}, nil
{{end}}default:
		return {{.Default}}, fmt.Errorf("%v(%T) can't convert to {{.Type}}", v, v)
	}
}
`

func buildTem(tem string, data interface{}) string {
	t, err := template.New("").Parse(tem)
	assert(err)

	buf := new(bytes.Buffer)
	assert(t.Execute(buf, data))

	return buf.String()
}

var temTest = `package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTo{{.TypeTitle}}(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		args       interface{}
		want       {{.Type}}
		wantErr    bool
		containErr string
	}{
{{range .TestCase}}
		{
			name: "{{.Name}}",
			args: {{if eq .InputType "str"}}"{{.Input}}"{{else}}{{.Input}}{{end}},
{{if .ContainErr}}
			wantErr: true,
			containErr: "{{.ContainErr}}",
{{else}}
			want: {{.Output}},
{{end}}
		},
{{end}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := To{{.TypeTitle}}(tt.args)
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
`

var temObjTo = `func (r *Object) To{{.TypeTitle}}() ({{.Type}}, error) {
	if r.err != nil {
		return {{.Default}}, r.err
	}
	return internal.To{{.TypeTitle}}(r.obj)
}
`

var temObjToTest = `
	t.Run("To{{.TypeTitle}}", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			res, err := lambda.New({{.SuccessTestCase.Input}}).To{{.TypeTitle}}()
			as.Nil(err)
			as.Equal({{.SuccessTestCase.Input}}, res)
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(anyVal).To{{.TypeTitle}}()
			as.NotNil(err)
			as.Equal("{}(lambda_test.item) can't convert to {{.Type}}", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, err := lambda.New(nil).WithErr(fmt.Errorf("er")).To{{.TypeTitle}}()
			as.NotNil(err)
		})
	})
`
