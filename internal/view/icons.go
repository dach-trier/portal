package view

import (
	"bytes"
	"html/template"
)

type Icons template.Template

func NewIcons() *Icons {
	var tmpl *template.Template

	tmpl = template.New("")
	tmpl = template.Must(tmpl.ParseGlob("icons/*.html"))

	return (*Icons)(tmpl)
}

func (icons *Icons) FuncMap() template.FuncMap {
	return template.FuncMap{
		"icon": func(id string, attributes template.HTMLAttr) (template.HTML, error) {
			var buf bytes.Buffer
			if err := (*template.Template)(icons).ExecuteTemplate(&buf, id, attributes); err != nil {
				return "", err
			}
			return template.HTML(buf.String()), nil
		},
	}
}
