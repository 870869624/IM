package api

import (
	"fmt"
	"net/http"
	"strconv"
	db "wechat/db/sqlc"
	"wechat/util"

	"github.com/gin-gonic/gin"
)

// 添加朋友（应该是一个事物操作，修改还有请求状态，然后加入好友列表中）
func (server *Server) CreateFriend(ctx *gin.Context) {
	var req db.CreateFriendParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	account := ctx.Request.Header.Get("account")
	if account != req.FUserAccount {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(fmt.Errorf("账户信息错误")))
		return
	}

	//修改申请状态
	find := &db.GetApplyParams{
		ApplicateAccount: req.FUserAccount,
		TargetAccount:    req.TUserAccount,
		Object:           0,
	}
	applies, err := server.store.GetApply(ctx, *find)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	if len(applies) == 0 {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(fmt.Errorf("没有申请记录")))
		return
	}

	for _, v := range applies {
		v.Status = 1
		v.AType = 3
		arg := &db.CreateApplyParams{
			ApplicateAccount: v.ApplicateAccount,
			TargetAccount:    v.TargetAccount,
			AType:            v.AType,
			Status:           v.Status,
			Object:           v.Object,
		}
		//存入新的申请状态的
		_, err = server.store.CreateApply(ctx, *arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
			return
		}
		delete := &db.DeleteApplyParams{
			ApplicateAccount: v.ApplicateAccount,
			TargetAccount:    v.TargetAccount,
			Status:           0,
		}
		//删除已经修改过的
		err = server.store.DeleteApply(ctx, *delete)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
			return
		}
	}

	//将朋友加入好友列表
	friend, err := server.store.CreateFriend(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	fmt.Println(friend)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})

}

// 获取好友列表（对于单个用户来讲， 但是实际需要查询是否在对方的好友列表--在加好友的操作时）
type ListFriendReq struct {
	Fromaccount string `form:"fromaccount"`
	PageID      string `form:"pageid"`
	PageSize    string `form:"pagesize"`
}

func (server *Server) ListFriend(ctx *gin.Context) {
	var req ListFriendReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}
	account := ctx.Request.Header.Get("account")
	if account != req.Fromaccount {
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

	arg := &db.ListFriendParams{
		FUserAccount: req.Fromaccount,
		Limit:        int32(pagesize),
		Offset:       (int32(pageid-1) * int32(pagesize)),
	}
	friends, err := server.store.ListFriend(ctx, *arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"yourfriend": friends,
	})
}

//获取单个好友信息

//删除好友

//拒绝添加好友
