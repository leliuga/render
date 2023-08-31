package html

import (
	"sync"
	"time"

	"github.com/flosch/pongo2/v6"
	"github.com/gofiber/fiber/v2"
)

type (
	// Config represents a HTML renderer configuration.
	Config struct {
		Directory string                           `json:"directory" yaml:"Directory"`
		Debug     bool                             `json:"debug"     yaml:"Debug"`
		Minify    bool                             `json:"minify"    yaml:"Minify"`
		Cache     bool                             `json:"cache"     yaml:"Cache"`
		Variables pongo2.Context                   `json:"variables" yaml:"Variables"`
		filters   map[string]pongo2.FilterFunction `json:"-"`
	}

	// Html represents a HTML renderer.
	Html struct {
		Blocks   Blocks
		Partials Partials
		Layouts  Layouts
		Pages    Pages
		config   *Config
		loaded   bool
		storage  fiber.Storage
		mutex    *sync.Mutex
	}

	// Block represents a block in a template.
	Block struct {
		template *pongo2.Template
		minify   bool
	}

	// Partial represents a partial in a template.
	Partial struct {
		template *pongo2.Template
		minify   bool
	}

	// Layout represents a layout in a template.
	Layout struct {
		Variables pongo2.Context `yaml:"Variables"`
		template  *pongo2.Template
		minify    bool
	}

	// Page represents a page in a template.
	Page struct {
		Title                  string         `yaml:"Title"`
		Description            string         `yaml:"Description"`
		Keywords               []string       `yaml:"Keywords"`
		Language               string         `yaml:"Language"`
		Path                   string         `yaml:"Path"`
		Layout                 string         `yaml:"Layout"`
		Author                 string         `yaml:"Author"`
		RobotsIndex            bool           `yaml:"RobotsIndex"`
		RobotsFollow           bool           `yaml:"RobotsFollow"`
		Draft                  bool           `yaml:"Draft"`
		Static                 bool           `yaml:"Static"`
		SitemapChangeFrequency string         `yaml:"SitemapChangeFrequency"`
		SitemapPriority        float64        `yaml:"SitemapPriority"`
		CreatedAt              time.Time      `yaml:"CreatedAt"`
		UpdatedAt              time.Time      `yaml:"UpdatedAt"`
		Variables              pongo2.Context `yaml:"Variables"`
		template               *pongo2.Template
		minify                 bool
	}

	Blocks   map[string]*Block
	Partials map[string]*Partial
	Layouts  map[string]*Layout
	Pages    map[string]*Page
)
