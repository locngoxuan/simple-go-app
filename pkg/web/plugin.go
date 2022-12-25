package web

import (
	"simpleapp/pkg/def"
	"simpleapp/plugin"

	"github.com/urfave/cli/v2"
)

type HTTPOption struct {
}

type HTTPPlugin struct {
	plugin.EmptyPlugin
	HTTPOption
}

func NewHTTP(ctx *cli.Context) def.HTTP {
	opt := HTTPOption{}
	return &HTTPPlugin{
		HTTPOption: opt,
	}
}

func (h *HTTPPlugin) Start() error {
	return nil
}

func (h *HTTPPlugin) Stop() error {
	return nil
}
