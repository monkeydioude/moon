package template

import (
	"bytes"
	"html/template"
)

type Template interface {
	GetPath() string
	GetChilds() map[string]Template
	GetValues() interface{}
	Render() []byte
}

type Values interface {
	Get(string) []byte
	IsEmpty() bool
	GetAll() interface{}
}

type Engine struct {
	Template
	Path   string
	Childs map[string]Template
	Values interface{}
}

func NewEngine(path string) *Engine {
	return &Engine{
		Path:   path,
		Childs: make(map[string]Template),
	}
}

func (t *Engine) GetPath() string {
	return t.Path
}

func (t *Engine) GetChilds() map[string]Template {
	return t.Childs
}

func (t *Engine) GetValues() interface{} {
	return t.Values
}

func (t *Engine) WithChild(path, name string) *Engine {
	t.Childs[name] = NewEngine(path)
	return t.Childs[name].(*Engine)
}

func (t *Engine) Render() []byte {
	b := &bytes.Buffer{}
	tmpl, _ := template.ParseFiles(t.GetPath())
	childs := t.GetChilds()

	if len(childs) > 0 {
		childsContent := make(map[string]template.HTML)
		for n, tmp := range childs {
			childsContent[n] = template.HTML(tmp.Render())
		}
		tmpl.Execute(b, childsContent)
	}

	if t.GetValues() != nil {
		tmpl.Execute(b, t.GetValues())
	}

	return b.Bytes()
}

func (t *Engine) AddValue(v interface{}) *Engine {
	t.Values = v
	return t
}

type TemplateValues struct {
	values map[string][]byte
}

func (v *TemplateValues) Get(k string) []byte {
	val, ok := v.values[k]

	if !ok {
		return nil
	}

	return val
}

func (v *TemplateValues) IsEmpty() bool {
	return len(v.values) == 0
}
