package router

import (
	"github.com/gin-gonic/gin"
	v1 "wss/controller/v1"

	"fmt"
	"github.com/unrolled/secure"
	"wss/lib/setting"
	"wss/middleware/app_middleware"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	//r.Use(TlsHandler())
	//r.RunTLS(fmt.Sprintf(":%d", setting.ServerSetting.HttpsPort), setting.ServerSetting.CertFile, setting.ServerSetting.KeyFile)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(app_middleware.App())

	apiv1 := r.Group("/api/v1")
	apiv1.POST("/push_mqtt_msg", v1.PushMqttMsg)
	return r
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:" + fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
