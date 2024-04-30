package api

import (
	db "wechat/db/sqlc"
	"wechat/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	// tokenmaker string
	router *gin.Engine
}

func NewServer(store *db.Store) (*Server, error) {
	server := &Server{store: store}
	server.setunpRouter()
	return server, nil
}

func (server *Server) setunpRouter() {
	router := gin.Default()

	router.POST("/signin", server.CreateUser)
	router.POST("/login", server.loginUser)
	userWithAuth := router.Group("/users")
	userWithAuth.Use(middleware.ParaseToken())
	{
		userWithAuth.PATCH("/updata")
		userWithAuth.DELETE("/logout")
	}

	server.router = router

}

func (server *Server) Strat(address string) {
	server.router.Run(address)
}
