-- name: ListTranslatedInitiativesWithThumbnail :many
SELECT
    i.id,
    i.kind,
    t.name,
    t.description,
    img.id  AS image_id,
    img.url AS image_url
FROM initiatives i
LEFT JOIN initiative_translations t
    ON  t.initiative_id = i.id
    AND t.lang = $1
LEFT JOIN LATERAL (
    SELECT im.id, im.url
    FROM initiative_images ii
    JOIN images im ON im.id = ii.image_id
    WHERE ii.initiative_id = i.id
    LIMIT 1
) img ON true
WHERE (sqlc.arg(kind) ::text IS NULL OR i.kind = sqlc.arg(kind))
AND   (sqlc.arg(after)::text IS NULL OR i.id   > sqlc.arg(after))
ORDER BY i.id ASC
LIMIT $2;
