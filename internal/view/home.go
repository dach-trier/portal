package view

import (
	"html/template"
	"io"

	"golang.org/x/text/language"

	"github.com/dach-trier/i18n"
	"github.com/dach-trier/i18n/html"
	"github.com/dach-trier/portal/internal/model"
	"github.com/dach-trier/template/html"
)

type Home template.Template

func NewHome(icons *Icons, localization i18n.Bundle) *Home {
	var tmpl *template.Template

	tmpl = template.New("")
	tmpl = tmpl.Funcs(html.FuncMap())
	tmpl = tmpl.Funcs(icons.FuncMap())
	tmpl = tmpl.Funcs(i18n_html.FuncMap(localization))
	tmpl = template.Must(
		tmpl.ParseFiles(
			"templates/graphics/dach-logo-simplified.html",
			"templates/components/email-address.html",
			"templates/layouts/base.html",
			"templates/pages/home.html",
		),
	)

	return (*Home)(tmpl)
}

func (home *Home) RenderPage(
	w io.Writer,
	lang language.Tag,
	initiatives []model.TranslatedInitiativeWithThumbnail,
) error {
	return (*template.Template)(home).ExecuteTemplate(w, "layout.base", map[string]any{
		"Title":       "DACH e.V. Trier",
		"Lang":        lang,
		"Initiatives": initiatives,
	})
}
