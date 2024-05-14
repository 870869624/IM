-- name: CreateApply :one
INSERT INTO applies (
  applicate_account, target_account,a_type,status, object
) VALUES (
  $1, $2,$3, $4, $5
)
RETURNING *;

-- name: ListSendApply :many
SELECT * FROM applies
WHERE applicate_account = $1
ORDER BY target_account LIMIT $2 OFFSET $3;

-- name: ListReceivedApply :many
SELECT * FROM applies
WHERE target_account = $1
ORDER BY applicate_account LIMIT $2 OFFSET $3;

-- name: GetApply :many
SELECT * FROM applies
WHERE target_account = $1 and applicate_account = $2 and object = $3;

-- name: DeleteApply :exec
DELETE FROM applies WHERE applicate_account = $1 and target_account = $2 and status = $3;

