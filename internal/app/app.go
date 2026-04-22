package app

import (
	"bytes"
	"context"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/text/language"

	"github.com/dach-trier/i18n"
	"github.com/dach-trier/i18n/http"
	"github.com/dach-trier/portal/internal/query"
	"github.com/dach-trier/portal/internal/repo"
	"github.com/dach-trier/portal/internal/view"

	chi "github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

type App struct {
	localization i18n.Bundle
	repos        repo.Bundle
	views        struct {
		initiatives *view.InitiativeListing
	}
}

func New(repos repo.Bundle) *App {
	app := &App{}

	app.localization = i18n.NewBundle()
	app.localization.MustLoadYaml(os.DirFS("i18n"), "de.yml", language.German)
	app.localization.MustLoadYaml(os.DirFS("i18n"), "uk.yml", language.Ukrainian)

	app.repos = repos

	// --

	icons := view.NewIcons()
	app.views.initiatives = view.NewInitiativeListing(icons, app.localization)

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
	router.Get("/initiatives", app.initiatives)

	return router
}

func (app *App) initiatives(w http.ResponseWriter, r *http.Request) {
	html := bytes.NewBuffer(make([]byte, 0, 1024))
	lang := r.Context().Value("lang").(language.Tag)
	kind := r.URL.Query().Get("kind")
	// after := r.URL.Query().Get("after")

	// --
	// load initiatives
	// --

	filter := query.InitiativeFilter{Kind: kind}
	cursor := query.Cursor[string]{Limit: math.MaxInt32}

	initiatives, err := app.repos.Initiatives.ListTranslatedInitiativesWithThumbnail(
		context.Background(),
		lang,
		filter,
		cursor,
	)

	if err != nil {
		http.Error(w, "failed to load initiatives", http.StatusInternalServerError)
		log.Printf("failed to load initiatives\n")
		log.Printf("reason: %v\n", err)
		return
	}

	// --
	// render
	// --

	err = app.views.initiatives.RenderPage(html, lang, initiatives)

	if err != nil {
		http.Error(w, "failed to render initiatives", http.StatusInternalServerError)
		log.Printf("failed to render initiatives\n")
		log.Printf("reason: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(html.Len()))
	w.WriteHeader(http.StatusOK)

	_, err = html.WriteTo(w)

	if err != nil {
		log.Printf("failed writing response\n")
		log.Printf("reason: %v\n", err)
		return
	}
}
