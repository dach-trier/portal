package sqlc_repo

import (
	"context"

	"golang.org/x/text/language"

	"github.com/dach-trier/portal/database/sqlc"
	"github.com/dach-trier/portal/internal/model"
	"github.com/dach-trier/portal/internal/repo"
)

type InitiativeRepository sqlc.Queries

func NewInitiativeRepository(db sqlc.DBTX) *InitiativeRepository {
	return (*InitiativeRepository)(sqlc.New(db))
}

func (repo *InitiativeRepository) ListTranslatedInitiativesWithThumbnail(
	ctx context.Context,
	lang language.Tag,
	filter repo.InitiativeFilter,
	pagination repo.Pagination,
) ([]model.TranslatedInitiativeWithThumbnail, error) {
	queries := (*sqlc.Queries)(repo)

	result, err := queries.ListTranslatedInitiativesWithThumbnail(
		ctx,
		sqlc.ListTranslatedInitiativesWithThumbnailParams{
			Lang:   lang.String(),
			Kind:   filter.Kind,
			Limit:  int32(pagination.Limit),
			Offset: int32(pagination.Offset),
		},
	)

	if err != nil {
		return nil, err
	}

	initiatives := []model.TranslatedInitiativeWithThumbnail{}

	for _, row := range result {
		initiatives = append(
			initiatives,
			model.TranslatedInitiativeWithThumbnail{
				ID:          row.ID,
				Kind:        row.Kind,
				Name:        row.Name.String,
				Description: row.Description.String,
				ImageID:     row.ImageID,
				ImageUrl:    row.ImageUrl,
			},
		)
	}

	return initiatives, nil
}
