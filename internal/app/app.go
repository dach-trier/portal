package app

import (
	"net/http"
	"os"

	"golang.org/x/text/language"

	"github.com/dach-trier/i18n"
	"github.com/dach-trier/i18n/http"

	chi "github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

type App struct {
	localization i18n.Bundle
}

func New() *App {
	app := &App{}

	app.localization = i18n.NewBundle()
	app.localization.MustLoadYaml(os.DirFS("i18n"), "de.yml", language.German)
	app.localization.MustLoadYaml(os.DirFS("i18n"), "uk.yml", language.Ukrainian)

	// --

	return app
}

func (app *App) Router() http.Handler {
	assets := http.FileServer(http.Dir("assets"))

	router := chi.NewRouter()
	router.Use(chi_middleware.Logger)
	router.Use(i18n_http.Middleware(func(tags ...language.Tag) language.Tag {
		fallback := language.English

		for _, tag := range tags {
			switch tag {
			case language.AmericanEnglish, language.BritishEnglish, language.English:
				return language.English

			case language.German, language.Ukrainian:
				return tag

			// A significant number of Ukrainian-speaking users may have
			// Russian set as their primary language. Since Russian isn't
			// supported, we simply use Ukrainian translations instead.
			case language.Russian:
				fallback = language.Ukrainian
			}
		}

		return fallback

	}))

	router.Get("/assets/*", http.StripPrefix("/assets/", assets).ServeHTTP)

	return router
}
