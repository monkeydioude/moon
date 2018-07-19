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
	Get(string) template.HTML
	Add(string, string)
	IsEmpty() bool
	Values() map[string]template.HTML
}

// Engine implements Template
// This basic templeting engine allows to build a nested tree of templates for reuse
type Engine struct {
	Path      string
	Templates map[string]Template
	Values    *TemplateValues
}

// NewEngine produces an Engine
func NewEngine(path string) *Engine {
	return &Engine{
		Path:      path,
		Templates: make(map[string]Template),
		Values:    NewTemplateValues(),
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

func (t *Engine) AddValue(name, value string) *Engine {
	t.Values.Add(name, value)
	return t
}

type TemplateValues struct {
	values map[string]template.HTML
}

func NewTemplateValues() *TemplateValues {
	return &TemplateValues{
		values: make(map[string]template.HTML),
	}
}

func (v *TemplateValues) Get(k string) template.HTML {
	val, ok := v.values[k]

	if !ok {
		return ""
	}

	return val
}

func (v *TemplateValues) IsEmpty() bool {
	return len(v.values) == 0
}

func (v *TemplateValues) Add(name, value string) {
	v.values[name] = template.HTML(value)
}
func (v *TemplateValues) Values() map[string]template.HTML {
	return v.values
}

func (v *TemplateValues) Merge(from map[string]template.HTML) {
	for n, val := range from {
		v.values[n] = val
	}
}
