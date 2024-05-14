package tcp

import (
	db "wechat/db/sqlc"
)

type Server struct {
	store *db.Store
	// tokenmaker string
}

func NewServer(store *db.Store) (*Server, error) {
	server := &Server{store: store}
	return server, nil
}
