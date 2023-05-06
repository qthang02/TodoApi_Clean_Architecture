package middleware

import (
	"TodoApi/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInternal(err))
				}
			}
		}()

		c.Next()
	}
}
