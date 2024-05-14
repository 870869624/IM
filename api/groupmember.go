package api

import (
	"fmt"
	"net/http"
	"strconv"
	db "wechat/db/sqlc"
	"wechat/util"

	"github.com/gin-gonic/gin"
)

func (server *Server) JoinGroup(ctx *gin.Context) {
	var req db.JoinGroupParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	//修改申请状态
	find := &db.GetApplyParams{
		ApplicateAccount: req.UserAccount,
		TargetAccount:    req.GroupAccount,
		Object:           1,
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

	_, err = server.store.JoinGroup(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "成功加入",
	})

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
}

type ListGroupMembers struct {
	GroupAccount string `form:"groupaccount"`
	PageID       string `form:"pageid"`
	PageSize     string `form:"pagesize"`
}

func (server *Server) ListGroupMember(ctx *gin.Context) {
	var req ListGroupMembers
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
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
	arg := &db.ListGroupMemberParams{
		GroupAccount: req.GroupAccount,
		Limit:        int32(pagesize),
		Offset:       (int32(pageid) - 1) * int32(pagesize),
	}
	groupmembers, err := server.store.ListGroupMember(ctx, *arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"groupmembers": groupmembers,
	})
}

// 修改群内用户昵称
func (server *Server) ChangeUsername(ctx *gin.Context) {}
