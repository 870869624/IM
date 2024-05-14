-- name: CreateFriend :one
INSERT INTO friends (
  f_user_account, t_user_account
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetFriend :one
SELECT * FROM friends
where t_user_account = $1 and f_user_account = $2;

-- name: ListFriend :many
SELECT * FROM friends
where f_user_account = $1
ORDER BY t_user_account LIMIT $2 OFFSET $3;