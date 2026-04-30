-- name: ListProjects :many
SELECT p.id
FROM projects p
WHERE (p.id > sqlc.narg(after) OR sqlc.narg(after) IS NULL)
ORDER BY p.id ASC
LIMIT $1;

-- name: GetProjectTranslation :one
SELECT pt.name, pt.body
FROM project_translations pt
WHERE pt.project_id = $1 AND pt.lang = $2;

-- name: ListProjectAssets :many
SELECT a.id, a.type, a.url
FROM project_assets pa
JOIN assets a ON pa.asset_id = a.id
WHERE pa.project_id = $1
  AND (a.id   > sqlc.narg(after) OR sqlc.narg(after) IS NULL)
  AND (a.type = sqlc.narg(type)  OR sqlc.narg(type)  IS NULL)
ORDER BY a.id
LIMIT $2;
