package rfcv1

import "github.com/gin-gonic/gin"

type RESTFullControllerV1 struct {
	router *gin.RouterGroup

	PingPong *PingPongController
}

func NewRESTFullControllerV1(router *gin.RouterGroup) {
	controller := &RESTFullControllerV1{
		router:   router,
		PingPong: NewPingPongController(),
	}

	controller.init()
}

func (api *RESTFullControllerV1) init() {
	{
		router := api.router.Group("ping")
		router.GET("", api.PingPong.Get)
	}
}
