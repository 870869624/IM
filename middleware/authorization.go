package middleware

import (
	"fmt"
	"net/http"
	"wechat/util"

	"github.com/gin-gonic/gin"
)

func ParaseToken() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"token无效": "请登录",
			})
			ctx.Abort()
			return
		}
		fmt.Println(tokenString)

		claims, err := util.ParaseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrResponse(err))
			ctx.Abort()
			return
		}

		ctx.Request.Header.Set("account", claims.Account)
		ctx.Next()
	}
}
