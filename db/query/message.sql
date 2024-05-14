-- name: CreateMessage :one
INSERT INTO messages (
  content, from_user_account,to_user_account, group_account,m_type, networkstatus
) VALUES (
  $1, $2,$3,$4,$5, $6
)
RETURNING *;

-- name: GetMessageToUser :many
SELECT content, from_user_account, to_user_account, m_type FROM messages
WHERE to_user_account = $1;

-- name: GetMessageToGroup :many
SELECT content, from_user_account, group_account, m_type, networkstatus FROM messages
WHERE group_account = $1;

-- -- name: ListMessage :many
-- SELECT * FROM messages
-- ORDER BY account;

-- -- name: UpdateMessage :exec
-- UPDATE messages SET username = $2
-- WHERE id = $1;
