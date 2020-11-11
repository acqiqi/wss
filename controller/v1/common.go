package v1

import (
	"github.com/gin-gonic/gin"
	"wss/lib/e"
)

func PushMqttMsg(c *gin.Context) {
	e.ApiOk(c, "mdzz", e.GetEmptyStruct())
}
