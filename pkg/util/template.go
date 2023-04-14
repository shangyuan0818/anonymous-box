package util

import (
	"bytes"
	"text/template"
)

func RenderTemplate(raw string, tpl interface{}) (string, error) {
	t := template.New("template")

	t, err := t.Parse(raw)
	if err != nil {
		return "", err
	}

	buff := bytes.NewBuffer([]byte{})
	if err := t.ExecuteTemplate(buff, "template", tpl); err != nil {
		return "", err
	}

	return buff.String(), nil
}
