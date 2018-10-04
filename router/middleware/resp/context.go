package resp

import (
	"github.com/gin-gonic/gin"
)

const key = "resp"

func FromContext(c *gin.Context) *Resp {
	return &Resp{c}
}
