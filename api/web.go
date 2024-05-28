package api

import (
	"wechat/web"

	"github.com/gin-gonic/gin"
)

func (server *Server) websocket(ctx *gin.Context) {
	web.WsHandler(ctx.Writer, ctx.Request)
}
