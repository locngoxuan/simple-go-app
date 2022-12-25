package pkg

import (
	"simpleapp/pkg/def"
	"simpleapp/plugin"

	"github.com/rs/zerolog/log"
)

const (
	KeyHttp       = "pkg/http"
	KeyRepository = "pkg/repository"
)

func GetHttp() def.HTTP {
	var pkg def.HTTP
	err := plugin.Get(KeyHttp, &pkg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get http plugin")
	}
	return pkg
}

func GetRepository() def.Repository {
	var pkg def.Repository
	err := plugin.Get(KeyRepository, &pkg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get repository plugin")
	}
	return pkg
}
