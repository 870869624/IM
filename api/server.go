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
		userWithAuth.POST("/groupcreate", server.CreateGroup)
		userWithAuth.POST("/apply", server.CreateApply)
		userWithAuth.GET("/listreceivedapply/", server.ListReceivedApply)
		userWithAuth.GET("/listsendapply/", server.ListSendApply)
		userWithAuth.POST("/createfriend", server.CreateFriend)
		userWithAuth.GET("/listfriend/", server.ListFriend)
		userWithAuth.POST("/joingroup", server.JoinGroup)
		userWithAuth.GET("/listgroupmember/", server.ListGroupMember)
	}

	server.router = router

}

func (server *Server) Strat(address string) {
	server.router.Run(address)
}
