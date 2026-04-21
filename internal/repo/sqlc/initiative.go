package sqlc_repo

import (
	"context"
	"html/template"

	"golang.org/x/text/language"

	"github.com/dach-trier/portal/database/sqlc"
	"github.com/dach-trier/portal/internal/model"
	"github.com/dach-trier/portal/internal/query"
)

type InitiativeRepository sqlc.Queries

func NewInitiativeRepository(db sqlc.DBTX) *InitiativeRepository {
	return (*InitiativeRepository)(sqlc.New(db))
}

func (repo *InitiativeRepository) ListTranslatedInitiativesWithThumbnail(
	ctx context.Context,
	lang language.Tag,
	filter query.InitiativeFilter,
	cursor query.Cursor[string],
) ([]model.TranslatedInitiativeWithThumbnail, error) {
	queries := (*sqlc.Queries)(repo)
	initiatives := []model.TranslatedInitiativeWithThumbnail{}

	// --
	// translated initiatives
	// --

	params := sqlc.ListTranslatedInitiativesParams{}

	params.Lang = lang.String()
	params.Limit = int32(cursor.Limit)
	params.Kind = filter.Kind

	if cursor.After != nil {
		params.After.String = *cursor.After
		params.After.Valid = true
	}

	result, err := queries.ListTranslatedInitiatives(ctx, params)

	if err != nil {
		return nil, err
	}

	//

	for _, row := range result {
		initiative := model.TranslatedInitiativeWithThumbnail{}

		initiative.ID = row.ID
		initiative.Kind = row.Kind
		initiative.Name = template.HTML(row.Name.String)
		initiative.Description = template.HTML(row.Description.String)

		// --
		// thumbnail
		// --

		params := sqlc.ListInitiativeImagesParams{
			InitiativeID: row.ID,
			Limit:        1,
		}

		result, err := queries.ListInitiativeImages(ctx, params)

		if err != nil {
			return nil, err
		}

		if len(result) > 0 {
			initiative.Thumbnail =
				&model.Image{
					ID:  result[0].ID,
					Url: result[0].Url,
				}
		}

		initiatives = append(initiatives, initiative)
	}

	return initiatives, nil
}
