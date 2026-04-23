package view

import (
	"html/template"
	"io"

	"golang.org/x/text/language"

	"github.com/dach-trier/i18n"
	"github.com/dach-trier/i18n/html"
	"github.com/dach-trier/template/html"
)

type Gallery template.Template

func NewGallery(icons *Icons, bundle i18n.Bundle) *Gallery {
	var tmpl *template.Template

	tmpl = template.New("")
	tmpl = tmpl.Funcs(html.FuncMap())
	tmpl = tmpl.Funcs(icons.FuncMap())
	tmpl = tmpl.Funcs(i18n_html.FuncMap(bundle))
	tmpl = template.Must(
		tmpl.ParseFiles(
			"templates/graphics/dach-logo-simplified.html",
			"templates/components/email-address.html",
			"templates/layouts/base.html",
			"templates/pages/gallery.html",
		),
	)

	return (*Gallery)(tmpl)
}

func (gallery *Gallery) RenderPage(
	w io.Writer,
	lang language.Tag,
) error {
	return (*template.Template)(gallery).ExecuteTemplate(w, "layout.base", map[string]any{
		"Title": "DACH e.V. Trier",
		"Lang":  lang,
	})
}
