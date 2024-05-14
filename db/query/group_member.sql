-- name: JoinGroup :one
INSERT INTO group_members (
  name, group_account, user_account
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListGroupMember :many
SELECT * FROM group_members
where group_account = $1
ORDER BY user_account LIMIT $2 OFFSET $3;

-- name: GetGMAccount :many
SELECT user_account FROM group_members
where group_account = $1;

-- 获取用户所在的群账号
-- name: GetUserGroup :many
SELECT group_members.group_account FROM group_members
where group_members.user_account = $1;

-- name: CheckUser :one
SELECT user_account FROM group_members
where group_account = $1 and user_account = $2;