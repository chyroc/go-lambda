package internal

import (
	"bytes"
	"text/template"
)

func BuildTemplate(tem string, data interface{}) (string, error) {
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
