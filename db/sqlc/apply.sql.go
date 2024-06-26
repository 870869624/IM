// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: apply.sql

package db

import (
	"context"
)

const createApply = `-- name: CreateApply :one
INSERT INTO applies (
  applicate_account, target_account,a_type,status, object
) VALUES (
  $1, $2,$3, $4, $5
)
RETURNING id, applicate_account, target_account, a_type, status, object
`

type CreateApplyParams struct {
	ApplicateAccount string `json:"applicate_account"`
	TargetAccount    string `json:"target_account"`
	AType            int32  `json:"a_type"`
	Status           int32  `json:"status"`
	Object           int32  `json:"object"`
}

func (q *Queries) CreateApply(ctx context.Context, arg CreateApplyParams) (Apply, error) {
	row := q.db.QueryRowContext(ctx, createApply,
		arg.ApplicateAccount,
		arg.TargetAccount,
		arg.AType,
		arg.Status,
		arg.Object,
	)
	var i Apply
	err := row.Scan(
		&i.ID,
		&i.ApplicateAccount,
		&i.TargetAccount,
		&i.AType,
		&i.Status,
		&i.Object,
	)
	return i, err
}

const deleteApply = `-- name: DeleteApply :exec
DELETE FROM applies WHERE applicate_account = $1 and target_account = $2 and status = $3
`

type DeleteApplyParams struct {
	ApplicateAccount string `json:"applicate_account"`
	TargetAccount    string `json:"target_account"`
	Status           int32  `json:"status"`
}

func (q *Queries) DeleteApply(ctx context.Context, arg DeleteApplyParams) error {
	_, err := q.db.ExecContext(ctx, deleteApply, arg.ApplicateAccount, arg.TargetAccount, arg.Status)
	return err
}

const getApply = `-- name: GetApply :many
SELECT id, applicate_account, target_account, a_type, status, object FROM applies
WHERE target_account = $1 and applicate_account = $2 and object = $3
`

type GetApplyParams struct {
	TargetAccount    string `json:"target_account"`
	ApplicateAccount string `json:"applicate_account"`
	Object           int32  `json:"object"`
}

func (q *Queries) GetApply(ctx context.Context, arg GetApplyParams) ([]Apply, error) {
	rows, err := q.db.QueryContext(ctx, getApply, arg.TargetAccount, arg.ApplicateAccount, arg.Object)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Apply{}
	for rows.Next() {
		var i Apply
		if err := rows.Scan(
			&i.ID,
			&i.ApplicateAccount,
			&i.TargetAccount,
			&i.AType,
			&i.Status,
			&i.Object,
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

const listReceivedApply = `-- name: ListReceivedApply :many
SELECT id, applicate_account, target_account, a_type, status, object FROM applies
WHERE target_account = $1
ORDER BY applicate_account LIMIT $2 OFFSET $3
`

type ListReceivedApplyParams struct {
	TargetAccount string `json:"target_account"`
	Limit         int32  `json:"limit"`
	Offset        int32  `json:"offset"`
}

func (q *Queries) ListReceivedApply(ctx context.Context, arg ListReceivedApplyParams) ([]Apply, error) {
	rows, err := q.db.QueryContext(ctx, listReceivedApply, arg.TargetAccount, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Apply{}
	for rows.Next() {
		var i Apply
		if err := rows.Scan(
			&i.ID,
			&i.ApplicateAccount,
			&i.TargetAccount,
			&i.AType,
			&i.Status,
			&i.Object,
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

const listSendApply = `-- name: ListSendApply :many
SELECT id, applicate_account, target_account, a_type, status, object FROM applies
WHERE applicate_account = $1
ORDER BY target_account LIMIT $2 OFFSET $3
`

type ListSendApplyParams struct {
	ApplicateAccount string `json:"applicate_account"`
	Limit            int32  `json:"limit"`
	Offset           int32  `json:"offset"`
}

func (q *Queries) ListSendApply(ctx context.Context, arg ListSendApplyParams) ([]Apply, error) {
	rows, err := q.db.QueryContext(ctx, listSendApply, arg.ApplicateAccount, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Apply{}
	for rows.Next() {
		var i Apply
		if err := rows.Scan(
			&i.ID,
			&i.ApplicateAccount,
			&i.TargetAccount,
			&i.AType,
			&i.Status,
			&i.Object,
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
