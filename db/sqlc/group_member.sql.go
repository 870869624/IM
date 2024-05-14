// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: group_member.sql

package db

import (
	"context"
)

const checkUser = `-- name: CheckUser :one
SELECT user_account FROM group_members
where group_account = $1 and user_account = $2
`

type CheckUserParams struct {
	GroupAccount string `json:"group_account"`
	UserAccount  string `json:"user_account"`
}

func (q *Queries) CheckUser(ctx context.Context, arg CheckUserParams) (string, error) {
	row := q.db.QueryRowContext(ctx, checkUser, arg.GroupAccount, arg.UserAccount)
	var user_account string
	err := row.Scan(&user_account)
	return user_account, err
}

const getGMAccount = `-- name: GetGMAccount :many
SELECT user_account FROM group_members
where group_account = $1
`

func (q *Queries) GetGMAccount(ctx context.Context, groupAccount string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getGMAccount, groupAccount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var user_account string
		if err := rows.Scan(&user_account); err != nil {
			return nil, err
		}
		items = append(items, user_account)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserGroup = `-- name: GetUserGroup :many
SELECT group_members.group_account FROM group_members
where group_members.user_account = $1
`

// 获取用户所在的群账号
func (q *Queries) GetUserGroup(ctx context.Context, userAccount string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getUserGroup, userAccount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var group_account string
		if err := rows.Scan(&group_account); err != nil {
			return nil, err
		}
		items = append(items, group_account)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const joinGroup = `-- name: JoinGroup :one
INSERT INTO group_members (
  name, group_account, user_account
) VALUES (
  $1, $2, $3
)
RETURNING id, name, group_account, user_account, created_at
`

type JoinGroupParams struct {
	Name         string `json:"name"`
	GroupAccount string `json:"group_account"`
	UserAccount  string `json:"user_account"`
}

func (q *Queries) JoinGroup(ctx context.Context, arg JoinGroupParams) (GroupMember, error) {
	row := q.db.QueryRowContext(ctx, joinGroup, arg.Name, arg.GroupAccount, arg.UserAccount)
	var i GroupMember
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.GroupAccount,
		&i.UserAccount,
		&i.CreatedAt,
	)
	return i, err
}

const listGroupMember = `-- name: ListGroupMember :many
SELECT id, name, group_account, user_account, created_at FROM group_members
where group_account = $1
ORDER BY user_account LIMIT $2 OFFSET $3
`

type ListGroupMemberParams struct {
	GroupAccount string `json:"group_account"`
	Limit        int32  `json:"limit"`
	Offset       int32  `json:"offset"`
}

func (q *Queries) ListGroupMember(ctx context.Context, arg ListGroupMemberParams) ([]GroupMember, error) {
	rows, err := q.db.QueryContext(ctx, listGroupMember, arg.GroupAccount, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GroupMember{}
	for rows.Next() {
		var i GroupMember
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.GroupAccount,
			&i.UserAccount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}