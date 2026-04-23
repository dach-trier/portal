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

func (repo *ProjectRepository) ListLocalizedProjectsWithThumbnail(
	ctx context.Context,
	lang language.Tag,
	cursor query.Cursor[string],
) ([]model.LocalizedProjectWithThumbnail, error) {
	queries := (*sqlc.Queries)(repo)
	projects := []model.LocalizedProjectWithThumbnail{}

	// --
	// project ids
	// --

	listProjectsParams := sqlc.ListProjectsParams{}
	listProjectsParams.Limit = int32(cursor.Limit)

	if cursor.After != nil {
		listProjectsParams.After.String = *cursor.After
		listProjectsParams.After.Valid = true
	}

	result, err := queries.ListProjects(ctx, listProjectsParams)

	if err != nil {
		return nil, err
	}

	//

	for _, id := range result {
		project := model.LocalizedProjectWithThumbnail{ID: id}

		// --
		// localization
		// --

		getProjectLocalizationParams := sqlc.GetProjectLocalizationParams{}
		getProjectLocalizationParams.ProjectID = id
		getProjectLocalizationParams.Lang = lang.String()

		localization, err := queries.GetProjectLocalization(ctx, getProjectLocalizationParams)

		if err != nil {
			return nil, err
		}

		project.Name = template.HTML(localization.Name.String)
		project.Description = template.HTML(localization.Description.String)

		// --
		// thumbnail
		// --

		listProjectImagesParams := sqlc.ListProjectImagesParams{
			ProjectID: id,
			Limit:     1,
		}

		result, err := queries.ListProjectImages(ctx, listProjectImagesParams)

		if err != nil {
			return nil, err
		}

		if len(result) > 0 {
			project.Thumbnail =
				&model.Image{
					ID:  result[0].ID,
					Url: result[0].Url,
				}
		}

		projects = append(projects, project)
	}

	return projects, nil
}
