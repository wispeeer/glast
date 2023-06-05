package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/wispeeer/glast/internal/pkg/usage"
	"github.com/wispeeer/glast/pkg/info"
	"github.com/wispeeer/glast/pkg/system"
)

type app struct {
	Service func() error

	failed  int
	success int
}

func GetApp() *app {
	return &app{
		failed:  137,
		success: 0,
	}
}

func (app *app) Run() (int, error) {
	var err error
	var isBreak bool = false

	// check service instance
	if app.Service == nil {
		return app.failed, errors.New("exception: Service Instance NOT DEFINE")
	}

	// process app flags
	if len(os.Args) > 1 {
		var argv = os.Args[1:]
		var argc = len(os.Args) - 1
		isBreak, err = app.flags(argc, argv)
		if err != nil {
			return app.success, err
		}
	}

	// service entry
	if isBreak {
		return app.success, err
	}

	// run service
	err = app.Service()
	return app.success, err
}

func (app *app) flags(argc int, argv []string) (bool, error) {
	var err error = nil
	var isBreak bool = false

	switch argv[0] {
	case "-e", "--entry", "entry":
		isBreak = false
		if argc >= 3 {
			system.Entry = argv[1]
			system.GoEntry = argv[2]
		} else {
			err = fmt.Errorf("glast --entry ./ ./main.go")
		}
	case "-h", "--help", "help":
		isBreak = true
		usage.Usage("")
	case "-v", "--version", "version":
		isBreak = true
		info.Version.Print()
	default:
		isBreak = true
		err = errors.New("usage: glast -h")
	}
	return isBreak, err
}

var App = GetApp()
