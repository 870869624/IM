package api

import (
	"fmt"
	"net/http"
	"strconv"
	db "wechat/db/sqlc"
	"wechat/util"

	"github.com/gin-gonic/gin"
)

// 发送好友申请
func (server *Server) CreateApply(ctx *gin.Context) {
	var req db.CreateApplyParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	arg := &db.CreateApplyParams{
		ApplicateAccount: req.ApplicateAccount,
		TargetAccount:    req.TargetAccount,
		AType:            req.AType,
		Status:           req.Status,
		Object:           req.Object,
	}
	account := ctx.Request.Header.Get("account")
	if account != arg.ApplicateAccount {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(fmt.Errorf("账户信息错误")))
		return
	}

	group, err := server.store.CreateApply(ctx, *arg) //存入数据库
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	fmt.Println(group)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "成功发送申请",
	})
}

type ListReceivedApplyReq struct {
	TargetAccount string `form:"targetaccount"` //好友请求会发送到该目标账户
	PageID        string `form:"pageid"`
	PageSize      string `form:"pagesize"`
}

// 获取收到的好友请求记录
func (server *Server) ListReceivedApply(ctx *gin.Context) {
	var req ListReceivedApplyReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}
	account := ctx.Request.Header.Get("account")
	if account != req.TargetAccount {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(fmt.Errorf("账户信息错误")))
		return
	}

	//转换参数
	pageid, err := strconv.Atoi(req.PageID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}
	pagesize, err := strconv.Atoi(req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	arg := &db.ListReceivedApplyParams{
		TargetAccount: req.TargetAccount,
		Limit:         int32(pagesize),
		Offset:        (int32(pageid-1) * int32(pagesize)),
	}
	applys, err := server.store.ListReceivedApply(ctx, *arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	fmt.Println(applys)
	ctx.JSON(http.StatusOK, gin.H{
		"applytoyou": applys,
	})
}

type ListSendApplyReq struct {
	ApplicateAccount string `form:"applicateaccount"` //好友请求会发送到该目标账户
	PageID           string `form:"pageid"`
	PageSize         string `form:"pagesize"`
}

// 获取发送的的好友请求记录
func (server *Server) ListSendApply(ctx *gin.Context) {
	var req ListSendApplyReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}
	account := ctx.Request.Header.Get("account")
	if account != req.ApplicateAccount {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(fmt.Errorf("账户信息错误")))
		return
	}

	//转换参数
	pageid, err := strconv.Atoi(req.PageID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}
	pagesize, err := strconv.Atoi(req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	arg := &db.ListSendApplyParams{
		ApplicateAccount: req.ApplicateAccount,
		Limit:            int32(pagesize),
		Offset:           (int32(pageid-1) * int32(pagesize)),
	}
	applys, err := server.store.ListSendApply(ctx, *arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	fmt.Println(applys)
	ctx.JSON(http.StatusOK, gin.H{
		"yourapplies": applys,
	})
}
