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

type ProjectListing template.Template

func NewProjectListing(icons *Icons, bundle i18n.Bundle) *ProjectListing {
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
			"templates/pages/project-listing.html",
		),
	)

	return (*ProjectListing)(tmpl)
}

func (projectListing *ProjectListing) RenderPage(
	w io.Writer,
	lang language.Tag,
	projects []model.TranslatedProjectWithThumbnail,
) error {
	return (*template.Template)(projectListing).ExecuteTemplate(w, "layout.base", map[string]any{
		"Title":    "DACH e.V. Trier",
		"Lang":     lang,
		"Projects": projects,
	})
}
