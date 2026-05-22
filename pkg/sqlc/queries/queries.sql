-- name: GetOneUser :one
SELECT * FROM users
LIMIT 1;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1;

-- name: GetAllUser :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users
SET name = $2
WHERE email = $1;

-- name: AddUser :one
INSERT INTO users(
    name,email
) VALUES(
    $1,$2
)
RETURNING *;