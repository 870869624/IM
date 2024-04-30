package api

import (
	"fmt"
	"net/http"
	db "wechat/db/sqlc"
	"wechat/util"

	"github.com/gin-gonic/gin"
)

type GetUserRequest struct {
	Account        string `json:"account" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
}

// 用户注册接口
func (server *Server) CreateUser(ctx *gin.Context) {
	var userReq db.CreateUserParams

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}
	HashPassword, err := util.HashPassword(userReq.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       userReq.Username,
		HashedPassword: HashPassword,
		Phone:          userReq.Phone,
		Account:        userReq.Account,
	}
	userRsp, err := server.store.CreateUser(ctx, arg)
	fmt.Println(userRsp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, userRsp)
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrResponse(err))
		return
	}

	userRsp, err := server.store.GetUser(ctx, req.Account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	//密码不对
	err = util.CheckHashPassword(userRsp.HashedPassword, req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}
	tokenString, err := util.CreatToken(req.Account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tokenString": tokenString,
	})
}
