package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	*gin.Context
}

func (rc *Resp) BadRequest(msg string) {
	rc.Context.JSON(http.StatusBadRequest, gin.H{"code": 4000, "msg": msg})
}

func (rc *Resp) ServerUnavailable(msg string) {
	rc.Context.JSON(http.StatusServiceUnavailable, gin.H{"code": 5000, "msg": msg})
}

func (rc *Resp) OK(data interface{}) {
	rc.Context.JSON(http.StatusOK, gin.H{"code": 0, "data": data})
}
