package app

import (
	"fmt"

	"github.com/iivkis/vk-chunte/config"
	"github.com/iivkis/vk-chunte/internal/server/restfull"
)

func Launch() {
	// repo := repository.NewRespository()

	httpServer := restfull.NewRESTFullServer()

	if err := httpServer.Start(fmt.Sprintf("%s:%s", config.HOST, config.PORT)); err != nil {
		panic(err)
	}
}
