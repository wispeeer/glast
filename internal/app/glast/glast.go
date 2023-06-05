package glast

import (
	"github.com/wispeeer/glast/internal/app/glast/utils"
	"github.com/wispeeer/glast/pkg/logger"
)

func Service() error {
	var err error

	module := utils.ModuleName()
	logger.Task("init").Infof("module name: %s\n", module)

	return err
}
