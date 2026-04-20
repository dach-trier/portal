-- name: CreateInitiativeTranslation :one
INSERT INTO initiative_translations (initiative_id, lang, name, description)
VALUES ($1, $2, $3, $4)
RETURNING initiative_id, lang, name, description;

-- name: GetInitiativeTranslation :one
SELECT *
FROM initiative_translations
WHERE initiative_id = $1 AND lang = $2;

-- name: ListInitiativeTranslationsByInitiative :many
SELECT *
FROM initiative_translations
WHERE initiative_id = $1
ORDER BY lang
LIMIT $2
OFFSET $3;

-- name: UpdateInitiativeTranslation :one
UPDATE initiative_translations
SET name = $3,
    description = $4
WHERE initiative_id = $1 AND lang = $2
RETURNING initiative_id, lang, name, description;

-- name: DeleteInitiativeTranslation :exec
DELETE FROM initiative_translations
WHERE initiative_id = $1 AND lang = $2;
