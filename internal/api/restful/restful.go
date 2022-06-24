package restful

import (
	"github.com/gin-gonic/gin"
	restfulhv1 "github.com/iivkis/vk-chunte/internal/api/restful/v1"
	"github.com/iivkis/vk-chunte/internal/repository"
)

type RESTFul struct {
	engine *gin.Engine
}

func NewRESTFul(repo *repository.Repository) *RESTFul {
	engine := gin.Default()

	restfulhv1.NewRESTFulHandlerV1(engine.Group("/api/v1"), repo)

	return &RESTFul{
		engine: engine,
	}
}

func (h *RESTFul) RunServer(address string) error {
	return h.engine.Run(address)
}
