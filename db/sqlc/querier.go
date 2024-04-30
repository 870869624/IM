// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetMessageToGroup(ctx context.Context, groupID sql.NullString) ([]GetMessageToGroupRow, error)
	GetMessageToUser(ctx context.Context, toUserID sql.NullString) ([]GetMessageToUserRow, error)
	GetUser(ctx context.Context, account string) (User, error)
	ListUser(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)