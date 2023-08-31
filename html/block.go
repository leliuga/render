package html

import (
	"bytes"
	"fmt"
	"io"

	"github.com/flosch/pongo2/v6"
	"github.com/leliuga/render"
)

// NewBlock creates a new Block from the given content.
func NewBlock(set *pongo2.TemplateSet, name string, content []byte, minify bool) (b *Block, err error) {
	b = &Block{
		minify: minify,
	}
	if b.template, err = set.FromBytes(content); err != nil {
		return nil, fmt.Errorf("syntax error in the block '%s': content (%s)", name, err.Error())
	}

	return b, nil
}

// Render renders the block with the given variables.
func (b *Block) Render(writer io.Writer, variables pongo2.Context) error {
	if b.minify {
		var buf bytes.Buffer
		if err := b.template.ExecuteWriter(variables, &buf); err != nil {
			return err
		}

		minified, err := render.Minify("text/html", buf.Bytes())
		if err != nil {
			return err
		}

		_, err = writer.Write(minified)

		return nil
	}

	return b.template.ExecuteWriter(variables, writer)
}
