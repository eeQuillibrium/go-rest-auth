package auth

import (
	"html/template"
)

type TemplateWrapper struct {
	Tmpl *template.Template
}

func (t *TemplateWrapper) Initialize(pattern string) {
	t.Tmpl = template.Must(template.ParseGlob(pattern))
}
