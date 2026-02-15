-- name: CreateCharacter :one
INSERT INTO characters (
    account_id,
    name,
    level,
    experience,
    gold,
    guild_id,
    pos_x,
    pos_y,
    slot,
    class_id,
    status,
    equipment,
    inventory
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING *;

-- name: GetCharacterByID :one
SELECT *
FROM characters
WHERE id = $1
ORDER BY created_at DESC LIMIT 1;

-- name: GetCharactersByAccountID :many
SELECT *
FROM characters
WHERE account_id = $1
ORDER BY created_at DESC;

-- name: UpdateCharacter :exec
UPDATE characters
SET
    name = $2,
    level = $3,
    experience = $4,
    gold = $5,
    guild_id = $6,
    pos_x = $7,
    pos_y = $8,
    status = $9,
    equipment = $10,
    inventory = $11,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteCharacter :exec
DELETE FROM characters
WHERE id = $1;

-- name: GetCharacterByName :one
SELECT *
FROM characters
WHERE name = $1
ORDER BY created_at DESC LIMIT 1;
