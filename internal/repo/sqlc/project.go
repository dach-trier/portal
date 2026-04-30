package sqlc_repo

import (
	"context"
	"html/template"

	"golang.org/x/text/language"

	"github.com/dach-trier/portal/database/sqlc"
	"github.com/dach-trier/portal/internal/model"
	"github.com/dach-trier/portal/internal/query"
)

type ProjectRepository sqlc.Queries

func NewProjectRepository(db sqlc.DBTX) *ProjectRepository {
	return (*ProjectRepository)(sqlc.New(db))
}

func (repo *ProjectRepository) ListTranslatedProjectsWithThumbnail(
	ctx context.Context,
	lang language.Tag,
	cursor query.Cursor[int64],
) ([]model.TranslatedProjectWithThumbnail, error) {
	queries := (*sqlc.Queries)(repo)
	projects := []model.TranslatedProjectWithThumbnail{}

	// --
	// project ids
	// --

	listProjectsParams := sqlc.ListProjectsParams{}
	listProjectsParams.Limit = int32(cursor.Limit)

	if cursor.After != nil {
		listProjectsParams.After.Int64 = *cursor.After
		listProjectsParams.After.Valid = true
	}

	result, err := queries.ListProjects(ctx, listProjectsParams)

	if err != nil {
		return nil, err
	}

	//

	for _, id := range result {
		project := model.TranslatedProjectWithThumbnail{ID: id}

		// --
		// localization
		// --

		getProjectLocalizationParams := sqlc.GetProjectTranslationParams{}
		getProjectLocalizationParams.ProjectID = id
		getProjectLocalizationParams.Lang = lang.String()

		localization, err := queries.GetProjectTranslation(ctx, getProjectLocalizationParams)

		if err != nil {
			return nil, err
		}

		project.Name = template.HTML(localization.Name.String)
		project.Body = template.HTML(localization.Body.String)

		// --
		// thumbnail
		// --

		listProjectImagesParams := sqlc.ListProjectAssetsParams{}

		listProjectImagesParams.ProjectID = id
		listProjectImagesParams.Limit = 1
		listProjectImagesParams.Type.String = "image"
		listProjectImagesParams.Type.Valid = true

		result, err := queries.ListProjectAssets(ctx, listProjectImagesParams)

		if err != nil {
			return nil, err
		}

		if len(result) > 0 {
			project.Thumbnail =
				&model.Asset{
					ID:   result[0].ID,
					Type: result[0].Type,
					Url:  result[0].Url,
				}
		}

		projects = append(projects, project)
	}

	return projects, nil
}
