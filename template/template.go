package template

import (
	"bytes"
	"html/template"
	"log"
)

type Template interface {
	GetPath() string
	GetChilds() map[string]Template
	GetValues() interface{}
	Render() []byte
}

type Values interface {
	Get(string) []byte
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
		Values: &TemplateValues{},
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
	t = t.Childs[name].(*Engine)
	return t
}

func (t *Engine) Render() []byte {
	b := &bytes.Buffer{}
	RenderWithBuffer(t, b)
	return b.Bytes()
}

func RenderWithBuffer(t Template, b *bytes.Buffer) *bytes.Buffer {
	childs := t.GetChilds()

	tmpl, err := template.ParseFiles(t.GetPath())
	if childs != nil {
		childsContent := make(map[string]string)
		for n, tmp := range childs {
			b = RenderWithBuffer(tmp, b)
			childsContent[n] = b.String()
		}
		if err != nil {
			log.Printf("[ERR ] Could not parse file. Reason: %s", err)
			return b
		}

		// tmpl.Execute(b, childsContent)
	}
	tmpl.Execute(b, t.GetValues())

	return b
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
