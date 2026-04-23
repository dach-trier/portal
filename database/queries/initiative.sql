-- name: ListTranslatedInitiatives :many
SELECT
    i.id,
    t.name,
    t.description
FROM initiatives i
LEFT JOIN initiative_translations t
    ON  t.initiative_id = i.id
    AND t.lang = $1
WHERE (i.id > sqlc.narg(after) OR sqlc.narg(after) IS NULL)
ORDER BY i.id ASC
LIMIT $2;

-- name: ListInitiativeImages :many
SELECT
    img.id AS id,
    img.url AS url
FROM initiative_images ii
JOIN images img ON ii.image_id = img.id
WHERE ii.initiative_id = $1
  AND (img.id > sqlc.narg(after) OR sqlc.narg(after) IS NULL)
ORDER BY img.id
LIMIT $2;
