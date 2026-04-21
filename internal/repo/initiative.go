package repo

import (
	"context"

	"golang.org/x/text/language"

	"github.com/dach-trier/portal/internal/model"
	"github.com/dach-trier/portal/internal/query"
)

type InitiativeRepository interface {
	ListTranslatedInitiativesWithThumbnail(
		ctx context.Context,
		lang language.Tag,
		filter query.InitiativeFilter,
		cursor query.Cursor[string],
	) ([]model.TranslatedInitiativeWithThumbnail, error)
}
