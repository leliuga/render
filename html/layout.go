package html

import (
	"bytes"
	"fmt"
	"io"

	"github.com/flosch/pongo2/v6"
	"github.com/goccy/go-yaml"
	"github.com/leliuga/render"
)

// NewLayout creates a new Layout from the given content.
func NewLayout(set *pongo2.TemplateSet, name string, content []byte, minify bool) (l *Layout, err error) {
	l = &Layout{
		minify: minify,
	}
	parts := bytes.SplitN(content, []byte("=="), 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("syntax error in the layout '%s': front matter or content is missing", name)
	}

	if err = yaml.Unmarshal(bytes.TrimSpace(parts[0]), l); err != nil {
		return nil, fmt.Errorf("syntax error in the layout '%s': front matter (%s)", name, err.Error())
	}

	if l.template, err = set.FromBytes(bytes.TrimSpace(parts[1])); err != nil {
		return nil, fmt.Errorf("syntax error in the layout '%s': content (%s)", name, err.Error())
	}

	return l, nil
}

// Render renders the layout with the given page and variables.
func (l *Layout) Render(writer io.Writer, variables pongo2.Context) error {
	variables["layout"] = pongo2.Context{
		"variables": l.Variables,
	}

	if l.minify {
		var buf bytes.Buffer
		if err := l.template.ExecuteWriter(variables, &buf); err != nil {
			return err
		}

		minified, err := render.Minify("text/html", buf.Bytes())
		if err != nil {
			return err
		}

		_, err = writer.Write(minified)

		return nil
	}

	return l.template.ExecuteWriter(variables, writer)
}
