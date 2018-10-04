package pv

import (
	"github.com/gin-gonic/gin"
	"github.com/yicat/p2s/model"
	"github.com/yicat/p2s/router/middleware/resp"
	"github.com/yicat/p2s/store"
	"strconv"
)

func GetAppList(c *gin.Context) {
	appList, err := store.FromContext(c).GetAppList(0, 10)
	if err != nil {
		resp.FromContext(c).ServerUnavailable(err.Error())
		return
	}

	resp.FromContext(c).OK(appList)
}

func GetAppByID(c *gin.Context) {
	appID, err := strconv.Atoi(c.Param("appID"))
	if err != nil {
		resp.FromContext(c).BadRequest("appID 类型")
		return
	}
	app, err := store.FromContext(c).GetApp(int64(appID))
	if err != nil {
		resp.FromContext(c).ServerUnavailable("不存在此 app")
	}

	resp.FromContext(c).OK(app)
}

func CreateApp(c *gin.Context) {
	var req model.CreateAppReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.FromContext(c).BadRequest("")
		return
	}

	app := &model.App{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := store.FromContext(c).CreateApp(app); err != nil {
		resp.FromContext(c).ServerUnavailable(err.Error())
		return
	}

	resp.FromContext(c).OK(nil)
}

// 暂时不支持
func UpdateApp(c *gin.Context) {

}

// 暂时不支持权限管理与用户认证
func DeleteApp(c *gin.Context) {
	appID, err := strconv.Atoi(c.Param("appID"))
	if err != nil {
		resp.FromContext(c).ServerUnavailable("appID 类型错误")
		return
	}

	app := model.App{BaseModel: model.BaseModel{ID: int64(appID)}}
	if err := store.FromContext(c).DeleteApp(&app); err != nil {
		resp.FromContext(c).ServerUnavailable(err.Error())
		return
	}

	resp.FromContext(c).OK(nil)
}
