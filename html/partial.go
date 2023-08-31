package html

import (
	"bytes"
	"fmt"
	"io"

	"github.com/flosch/pongo2/v6"
	"github.com/leliuga/render"
)

// NewPartial creates a new Partial from the given content.
func NewPartial(set *pongo2.TemplateSet, name string, content []byte, minify bool) (p *Partial, err error) {
	p = &Partial{
		minify: minify,
	}
	if p.template, err = set.FromBytes(content); err != nil {
		return nil, fmt.Errorf("syntax error in the partial '%s': content (%s)", name, err.Error())
	}

	return p, nil
}

// Render renders the partial with the given variables.
func (p *Partial) Render(writer io.Writer, variables pongo2.Context) error {
	if p.minify {
		var buf bytes.Buffer
		if err := p.template.ExecuteWriter(variables, &buf); err != nil {
			return err
		}

		minified, err := render.Minify("text/html", buf.Bytes())
		if err != nil {
			return err
		}

		_, err = writer.Write(minified)

		return nil
	}
	return p.template.ExecuteWriter(variables, writer)
}
