package restfulhv1

import (
	"github.com/gin-gonic/gin"
	"github.com/iivkis/vk-chunte/internal/repository"
)

// all v1 controllers
type V1Controllers struct {
	PingPong *PingPongController
}

func NewV1Controllers() *V1Controllers {
	return &V1Controllers{
		PingPong: NewPingPongController(),
	}
}

// RESTFul API handler v1
type RESTFulHandlerV1 struct {
	router      *gin.RouterGroup
	controllers *V1Controllers
}

// Create RESTFul API handler v1
func NewRESTFulHandlerV1(router *gin.RouterGroup, repo *repository.Repository) {
	handler := &RESTFulHandlerV1{
		router:      router,
		controllers: NewV1Controllers(),
	}

	handler.init()
}

// Initializing routes
func (h *RESTFulHandlerV1) init() {
	{
		router := h.router.Group("ping")
		router.GET("", h.controllers.PingPong.Get)
	}
}
