package util

import "github.com/gin-gonic/gin"

func ErrResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
