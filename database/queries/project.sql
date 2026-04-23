-- name: ListProjects :many
SELECT p.id
FROM projects p
WHERE (p.id > sqlc.narg(after) OR sqlc.narg(after) IS NULL)
ORDER BY p.id ASC
LIMIT $1;

-- name: GetProjectLocalization :one
SELECT pl.name, pl.description
FROM project_localizations pl
WHERE pl.project_id = $1 AND pl.lang = $2;

-- name: ListProjectImages :many
SELECT img.id, img.url
FROM project_images pi
JOIN images img ON pi.image_id = img.id
WHERE pi.project_id = $1
  AND (img.id > sqlc.narg(after) OR sqlc.narg(after) IS NULL)
ORDER BY img.id
LIMIT $2;
