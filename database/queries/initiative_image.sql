-- name: AddImageToInitiative :one
INSERT INTO initiative_images (initiative_id, image_id)
VALUES ($1, $2)
RETURNING initiative_id, image_id;

-- name: RemoveImageFromInitiative :exec
DELETE FROM initiative_images
WHERE initiative_id = $1 AND image_id = $2;

-- name: ListImagesForInitiative :many
SELECT i.id, i.url
FROM images i
JOIN initiative_images ii
  ON i.id = ii.image_id
WHERE ii.initiative_id = $1
ORDER BY i.id
LIMIT $2
OFFSET $3;

-- name: ListInitiativesForImage :many
SELECT ini.id, ini.kind
FROM initiatives ini
JOIN initiative_images ii
  ON ini.id = ii.initiative_id
WHERE ii.image_id = $1
ORDER BY ini.id
LIMIT $2
OFFSET $3;
