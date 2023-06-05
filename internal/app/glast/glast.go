package glast

import (
	"github.com/wispeeer/glast/internal/app/glast/tools"
	"github.com/wispeeer/glast/internal/app/glast/utils"
	"github.com/wispeeer/glast/pkg/logger"
	"github.com/wispeeer/glast/pkg/system"
)

func Service() error {
	var err error

	entry := system.Entry
	goEntry := system.GoEntry

	if len(entry) == 0 {
		entry = "./"
	}
	if len(goEntry) == 0 {
		goEntry = "./main.go"
	}

	logger.Task("init").Infof("start entry: %s\n", entry)
	logger.Task("init").Infof("go module entry: %s\n", goEntry)
	module := utils.ModuleName(entry)
	logger.Task("init").Infof("module name: %s\n", module)

	tools.FileScan(entry)
	logger.Task("scan").Infof("file scan: %2d\n", (len(tools.GoFile)))

	tools.ASTParser()

	return err
}
