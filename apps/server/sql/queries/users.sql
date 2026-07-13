-- name: CreateUser :one
INSERT INTO users (
    username,
    phone_number,
    password_hash
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;


-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE phone_number = $1
LIMIT 1;