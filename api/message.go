package api

import (
	"net/http"
	db "wechat/db/sqlc"
	"wechat/util"

	"github.com/gin-gonic/gin"
)

func (server *Server) Createmessage(ctx *gin.Context) {

	var Req db.CreateMessageParams
	if err := ctx.ShouldBindJSON(&Req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}
}
