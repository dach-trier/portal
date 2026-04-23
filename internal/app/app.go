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
		home     *view.Home
		projects *view.ProjectListing
		gallery  *view.Gallery
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

	app.views.home = view.NewHome(icons, app.localization)
	app.views.projects = view.NewProjectListing(icons, app.localization)
	app.views.gallery = view.NewGallery(icons, app.localization)

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
	router.Get("/", app.home)
	router.Get("/projects", app.projects)
	router.Get("/gallery", app.gallery)

	return router
}

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	html := bytes.NewBuffer(make([]byte, 0, 1024))
	lang := r.Context().Value("lang").(language.Tag)

	// --
	// load projects
	// --

	projects, err := app.repos.Projects.ListLocalizedProjectsWithThumbnail(
		context.Background(),
		lang,
		query.Cursor[string]{Limit: math.MaxInt32},
	)

	if err != nil {
		http.Error(w, "failed to load projects", http.StatusInternalServerError)
		log.Printf("failed to load projects\n")
		log.Printf("reason: %v\n", err)
		return
	}

	// --
	// render
	// --

	err = app.views.home.RenderPage(html, lang, projects)

	if err != nil {
		http.Error(w, "failed to render home page", http.StatusInternalServerError)
		log.Printf("failed to render home page\n")
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

func (app *App) projects(w http.ResponseWriter, r *http.Request) {
	html := bytes.NewBuffer(make([]byte, 0, 1024))
	lang := r.Context().Value("lang").(language.Tag)
	// after := r.URL.Query().Get("after")

	// --
	// load projects
	// --

	projects, err := app.repos.Projects.ListLocalizedProjectsWithThumbnail(
		context.Background(),
		lang,
		query.Cursor[string]{Limit: math.MaxInt32},
	)

	if err != nil {
		http.Error(w, "failed to load projects", http.StatusInternalServerError)
		log.Printf("failed to load projects\n")
		log.Printf("reason: %v\n", err)
		return
	}

	// --
	// render
	// --

	err = app.views.projects.RenderPage(html, lang, projects)

	if err != nil {
		http.Error(w, "failed to render projects", http.StatusInternalServerError)
		log.Printf("failed to render projects\n")
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

func (app *App) gallery(w http.ResponseWriter, r *http.Request) {
	html := bytes.NewBuffer(make([]byte, 0, 1024))
	lang := r.Context().Value("lang").(language.Tag)

	err := app.views.gallery.RenderPage(html, lang)

	if err != nil {
		http.Error(w, "failed to render projects", http.StatusInternalServerError)
		log.Printf("failed to render projects\n")
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
