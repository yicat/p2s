package pv

import (
	"github.com/gin-gonic/gin"
	"github.com/yicat/p2s/model"
	"github.com/yicat/p2s/router/middleware/resp"
	"github.com/yicat/p2s/store"
)

func GetAppList(c *gin.Context) {
	appList, err := store.FromContext(c).GetAppList(0, 10)
	if err != nil {
		resp.FromContext(c).ServerUnavailableWithMsg(err.Error())
		return
	}

	resp.FromContext(c).OKWithJSON(appList)
}

func GetAppByID(c *gin.Context) {

}

func CreateApp(c *gin.Context) {
	var req model.CreateAppReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.FromContext(c).BadRequest()
		return
	}

	app := &model.App{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := store.FromContext(c).CreateApp(app); err != nil {
		resp.FromContext(c).ServerUnavailableWithMsg(err.Error())
		return
	}

	resp.FromContext(c).OK()
}

func UpdateApp(c *gin.Context) {

}

func DeleteApp(c *gin.Context) {

}
