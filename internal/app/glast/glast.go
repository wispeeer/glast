package glast

import (
	"github.com/wispeeer/glast/internal/app/glast/utils"
	"github.com/wispeeer/glast/pkg/logger"
	"github.com/wispeeer/glast/pkg/system"
)

func Service() error {
	var err error

	entry := system.Entry

	if len(entry) == 0 {
		entry = "./"
	}

	logger.Task("init").Infof("start entry: %s\n", entry)
	module := utils.ModuleName(entry)
	logger.Task("init").Infof("module name: %s\n", module)

	return err
}
