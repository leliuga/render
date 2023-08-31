package render

import (
	"errors"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

var (
	minifier = minify.New()
)

// Minify minifies the b based on the mimeType.
func Minify(mimeType string, content []byte) ([]byte, error) {
	m, err := minifier.Bytes(mimeType, content)
	if errors.Is(err, minify.ErrNotExist) {
		m = content
		err = nil
	}

	return m, err
}

func init() {
	minifier.Add("text/html", &html.Minifier{})
	minifier.Add("text/css", &css.Minifier{})
	minifier.Add("application/javascript", &js.Minifier{})
	minifier.Add("application/json", &json.Minifier{})
	minifier.Add("application/xml", &xml.Minifier{})
	minifier.Add("image/svg+xml", &svg.Minifier{})
}
