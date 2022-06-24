package rfcv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingPongController struct{}

func NewPingPongController() *PingPongController {
	return &PingPongController{}
}

func (c *PingPongController) Get(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
