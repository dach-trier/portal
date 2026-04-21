package repo

import (
	"context"

	"golang.org/x/text/language"

	"github.com/dach-trier/portal/internal/model"
)

type InitiativeFilter struct {
	Kind string
}

type InitiativeRepository interface {
	ListTranslatedInitiativesWithThumbnail(
		ctx context.Context,
		lang language.Tag,
		filter InitiativeFilter,
		cursor Cursor[string],
	) ([]model.TranslatedInitiativeWithThumbnail, error)
}
