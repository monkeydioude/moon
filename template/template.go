package template

import (
	"bytes"
	"html/template"
)

// Template is the base interface for templating with Moon
// Must hold base information for a template
type Template interface {
	GetPath() string
	GetBoundTemplates() map[string]Template
	GetValues() Values
	Render() []byte
}

// Values is the base interface for holding values inside a template
type Values interface {
	Add(string, string)
	IsEmpty() bool
	Values() map[string]template.HTML
}

// Engine implements Template
// This basic templeting engine allows to build a nested tree of templates for reuse
type Engine struct {
	Path      string
	Templates map[string]Template
	Values    HTML
}

// NewEngine produces an Engine
func NewEngine(path string) *Engine {
	return &Engine{
		Path:      path,
		Templates: make(map[string]Template),
		Values:    HTML{},
	}
}

func (t *Engine) GetPath() string {
	return t.Path
}

func (t *Engine) GetBoundTemplates() map[string]Template {
	return t.Templates
}

func (t *Engine) GetValues() Values {
	return t.Values
}

func (t *Engine) BindTemplate(path, name string) *Engine {
	t.Templates[name] = NewEngine(path)
	return t.Templates[name].(*Engine)
}

func (t *Engine) Render() []byte {
	b := &bytes.Buffer{}
	tmpl, _ := template.ParseFiles(t.GetPath())
	childs := t.GetBoundTemplates()
	content := make(map[string]template.HTML)

	if len(childs) > 0 {
		for n, tmp := range childs {
			content[n] = template.HTML(tmp.Render())
		}
	}

	t.Values.Merge(content)
	tmpl.Execute(b, t.GetValues().Values())

	return b.Bytes()
}

func (t *Engine) BindValue(name, value string) *Engine {
	t.Values.Add(name, value)
	return t
}

func (t *Engine) BindValues(values HTML) *Engine {
	t.Values.Merge(values)
	return t
}

type HTML map[string]template.HTML

func (h HTML) Add(n, v string) {
	h[n] = template.HTML(v)
}

func (h HTML) IsEmpty() bool {
	return len(h) == 0
}

func (h HTML) Values() map[string]template.HTML {
	return h
}

func (h HTML) Merge(v HTML) {
	for n, val := range v {
		h[n] = val
	}
}
