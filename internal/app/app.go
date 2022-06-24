package app

import (
	"fmt"

	"github.com/iivkis/vk-chunte/config"
	"github.com/iivkis/vk-chunte/internal/api/restful"
	"github.com/iivkis/vk-chunte/internal/repository"
)

func Launch() {
	repo := repository.NewRespository()
	httpServer := restful.NewRESTFul(repo)

	if err := httpServer.RunServer(fmt.Sprintf("%s:%s", config.HOST, config.PORT)); err != nil {
		panic(err)
	}
}
