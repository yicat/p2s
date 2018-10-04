package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yicat/p2s/pv"
	"github.com/yicat/p2s/router/middleware"
)

func Load() *gin.Engine {
	rootR := gin.New()

	rootR.Use(middleware.Store(true))

	appGroup := rootR.Group("/api/app")
	appGroup.GET("/", pv.GetAppList)
	appGroup.GET("/:appID", pv.GetAppByID)
	appGroup.POST("/", pv.CreateApp)
	appGroup.DELETE("/:appID", pv.DeleteApp)

	return rootR
}
