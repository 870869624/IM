-- name: CreateMessage :one
INSERT INTO messages (
  content, from_user_id,to_user_id, group_id,m_type
) VALUES (
  $1, $2,$3,$4,$5
)
RETURNING *;

-- name: GetMessageToUser :many
SELECT content, from_user_id, to_user_id, m_type FROM messages
WHERE to_user_id = $1;

-- name: GetMessageToGroup :many
SELECT content, from_user_id, group_id, m_type FROM messages
WHERE group_id = $1;

-- -- name: ListMessage :many
-- SELECT * FROM messages
-- ORDER BY account;

-- -- name: UpdateMessage :exec
-- UPDATE messages SET username = $2
-- WHERE id = $1;