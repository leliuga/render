package html

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/flosch/pongo2/v6"
	"github.com/goccy/go-yaml"
)

// NewPage creates a new Page from the given content.
func NewPage(set *pongo2.TemplateSet, name string, content []byte, minify bool) (p *Page, err error) {
	p = &Page{
		Language:     DefaultLanguage,
		Path:         "/" + strings.TrimPrefix(name, PagesPrefix),
		Layout:       DefaultLayout,
		RobotsIndex:  DefaultRobotsIndex,
		RobotsFollow: DefaultRobotsFollow,
		minify:       minify,
	}
	parts := bytes.SplitN(content, []byte("=="), 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("syntax error in the page '%s': front matter or content is missing", name)
	}

	if err = yaml.Unmarshal(bytes.TrimSpace(parts[0]), p); err != nil {
		return nil, fmt.Errorf("syntax error in the page '%s': front matter (%s)", name, err.Error())
	}

	if p.template, err = set.FromBytes(bytes.TrimSpace(parts[1])); err != nil {
		return nil, fmt.Errorf("syntax error in the page '%s': content (%s)", name, err.Error())
	}

	return p, nil
}

// Render renders the page with the given variables.
func (p *Page) Render(writer io.Writer, variables pongo2.Context) error {
	variables["page"] = pongo2.Context{
		"title":         p.Title,
		"description":   p.Description,
		"keywords":      strings.Join(p.Keywords, ","),
		"language":      p.Language,
		"path":          p.Path,
		"layout":        p.Layout,
		"author":        p.Author,
		"robots_index":  p.RobotsIndex,
		"robots_follow": p.RobotsFollow,
		"draft":         p.Draft,
		"static":        p.Static,
		"created_at":    p.CreatedAt.Format("2006-01-02 15:04:05"),
		"updated_at":    p.UpdatedAt.Format("2006-01-02 15:04:05"),
		"variables":     p.Variables,
	}
	return p.template.ExecuteWriter(variables, writer)
}
