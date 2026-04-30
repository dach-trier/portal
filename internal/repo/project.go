package repo

import (
	"context"

	"golang.org/x/text/language"

	"github.com/dach-trier/portal/internal/model"
	"github.com/dach-trier/portal/internal/query"
)

type ProjectRepository interface {
	ListTranslatedProjectsWithThumbnail(
		ctx context.Context,
		lang language.Tag,
		cursor query.Cursor[int64],
	) ([]model.TranslatedProjectWithThumbnail, error)
}
