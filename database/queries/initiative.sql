-- name: CreateInitiative :one
INSERT INTO initiatives (id, kind)
VALUES ($1, $2)
RETURNING id, kind;

-- name: GetInitiativeByID :one
SELECT *
FROM initiatives
WHERE id = $1;

-- name: ListInitiatives :many
SELECT *
FROM initiatives
LIMIT $1
OFFSET $2;

-- name: UpdateInitiative :one
UPDATE initiatives
SET kind = $2
WHERE id = $1
RETURNING id, kind;

-- name: DeleteInitiative :exec
DELETE FROM initiatives
WHERE id = $1;
