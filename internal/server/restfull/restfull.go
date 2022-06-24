package restfull

import (
	"github.com/gin-gonic/gin"
	rfcv1 "github.com/iivkis/vk-chunte/internal/server/restfull/v1"
)

type RESTFullServer struct {
	engine *gin.Engine
}

func NewRESTFullServer() *RESTFullServer {
	engine := gin.Default()
	rfcv1.NewRESTFullControllerV1(engine.Group("/api/v1"))
	return &RESTFullServer{engine: engine}
}

func (h *RESTFullServer) Start(address string) error {
	return h.engine.Run(address)
}
