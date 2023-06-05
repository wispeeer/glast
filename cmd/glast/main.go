package main

import (
	"os"

	"github.com/wispeeer/glast/internal/app"
	"github.com/wispeeer/glast/pkg/logger"
)

func main() {
	// bootstrap
	service := app.App
	service.Service = app.Cli.Service()
	code, err := service.Run()
	defer os.Exit(code)
	if err != nil {
		logger.Task("app").Error(err)
	}
}
