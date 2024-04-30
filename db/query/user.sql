-- name: CreateUser :one
INSERT INTO users (
  username, hashed_password,phone, account,token
) VALUES (
  $1, $2,$3,$4,$5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE account = $1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY account;

-- name: UpdateUser :exec
UPDATE users SET username = $2
WHERE id = $1;
