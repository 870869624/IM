-- name: CreateGroup :one
INSERT INTO groups (
  name, account,owner
) VALUES (
  $1, $2,$3
)
RETURNING *;

-- name: GetGroup :one
SELECT * FROM groups
WHERE account = $1;

