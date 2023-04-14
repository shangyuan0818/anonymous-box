package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	asserts := assert.New(t)

	type tpl struct {
		Name string
	}

	raw := "Hello, {{.Name}}"

	content, err := RenderTemplate(raw, &tpl{Name: "World"})
	asserts.NoError(err)
	asserts.Equal("Hello, World", content)
}
