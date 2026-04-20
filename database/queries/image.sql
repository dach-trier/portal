-- name: CreateImage :one
INSERT INTO images (url)
VALUES ($1)
RETURNING id, url;

-- name: GetImageByID :one
SELECT *
FROM images
WHERE id = $1;

-- name: ListImages :many
SELECT *
FROM images;

-- name: UpdateImage :one
UPDATE images
SET url = $2
WHERE id = $1
RETURNING id, url;

-- name: DeleteImage :exec
DELETE FROM images
WHERE id = $1;
