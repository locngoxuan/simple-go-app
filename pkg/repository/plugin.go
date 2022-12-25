package repository

import (
	"simpleapp/pkg/def"
	"simpleapp/plugin"

	"github.com/urfave/cli/v2"
)

type RepositoryOption struct {
}

type RepositoryPlugin struct {
	plugin.EmptyPlugin
	RepositoryOption
}

func NewRepository(ctx *cli.Context) def.Repository {
	opt := RepositoryOption{}
	return &RepositoryPlugin{
		RepositoryOption: opt,
	}
}

func (h *RepositoryPlugin) Start() error {
	return nil
}

func (h *RepositoryPlugin) Stop() error {
	return nil
}
