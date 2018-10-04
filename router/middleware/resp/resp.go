package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	*gin.Context
}

func (rc *Resp) BadRequestWithMsg(msg string) {
	rc.Context.JSON(http.StatusBadRequest, gin.H{"code": 4000, "msg": msg})
}

func (rc *Resp) BadRequest() {
	rc.BadRequestWithMsg("")
}

func (rc *Resp) ServerUnavailableWithMsg(msg string) {
	rc.Context.JSON(http.StatusServiceUnavailable, gin.H{"code": 5000, "msg": msg})
}

func (rc *Resp) ServerUnavailable() {
	rc.ServerUnavailableWithMsg("")
}

func (rc *Resp) OK() {
	rc.OKWithJSON(nil)
}

func (rc *Resp) OKWithJSON(data interface{}) {
	rc.Context.JSON(http.StatusOK, gin.H{"code": 0, "data": data})
}
