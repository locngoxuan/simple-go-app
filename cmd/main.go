package main

import (
	"os"
	"os/signal"
	"simpleapp/pkg"
	"simpleapp/pkg/repository"
	"simpleapp/pkg/web"
	"simpleapp/plugin"
	"strconv"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	version = "devel"
	binFile = "app"
)

func main() {
	app := &cli.App{
		Name:      binFile,
		Copyright: "Loc Ngo",
		Version:   version,
		Flags:     []cli.Flag{},
		Action:    run,
	}
	err := app.Run(os.Args)
	log.Info().Interface("args", os.Args).Msg("execute binary file with args")
	if err != nil {
		log.Error().Err(err).Msg("failed to run application")
	}
}

func setupZeroLog() error {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
	return nil
}

func run(ctx *cli.Context) error {
	err := plugin.Register(pkg.KeyHttp, web.NewHTTP(ctx))
	if err != nil {
		return err
	}
	err = plugin.Register(pkg.KeyRepository, repository.NewRepository(ctx))
	if err != nil {
		return err
	}
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGINT,
	)
	<-signalChannel
	signal.Stop(signalChannel)
	close(signalChannel)
	err = afterShutdown()
	if err != nil {
		log.Warn().Err(err).Msg("get error while shuting down application")
	}
	log.Info().Msg("application has been shutdown")
	return nil
}

func afterShutdown() error {
	log.Info().Msg("shutting down application...")
	plugin.Uninstall()
	return nil
}
