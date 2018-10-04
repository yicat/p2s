package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yicat/p2s/store"
)

func Store(debug bool) gin.HandlerFunc {
	// 初始化 DB
	store_ := store.New(debug)

	return func(c *gin.Context) {
		store.ToContext(c, store_)
		c.Next()
	}
}
