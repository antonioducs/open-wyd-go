-- name: CreateAccount :one
INSERT INTO accounts (username, password_hash)
VALUES ($1, $2)
RETURNING *;

-- name: GetAccountByUsername :one
SELECT *
FROM accounts
WHERE username = $1
ORDER BY created_at DESC LIMIT 1;

-- name: GetAccountByID :one
SELECT *
FROM accounts
WHERE id = $1
ORDER BY created_at DESC LIMIT 1;
