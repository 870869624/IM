package api

import (
	"fmt"
	"net/http"
	db "wechat/db/sqlc"
	"wechat/util"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateGroup(ctx *gin.Context) {
	var req db.CreateGroupParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	arg := db.CreateGroupParams{
		Name:    req.Name,
		Account: req.Account,
		Owner:   req.Owner,
	}
	account := ctx.Request.Header.Get("account")
	if account != arg.Owner {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(fmt.Errorf("账户信息错误")))
		return
	}

	_, err := server.store.CreateGroup(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	owner := &db.JoinGroupParams{
		Name:         req.Owner,
		GroupAccount: req.Account,
		UserAccount:  req.Owner,
	}
	_, err = server.store.JoinGroup(ctx, *owner)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}
