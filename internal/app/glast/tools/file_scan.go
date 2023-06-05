package tools

import (
	"os"
	"path/filepath"

	"github.com/wispeeer/glast/pkg/utils"
)

var GoFile []string

func FileScan(start string) {
	if utils.IsExist(start) {
		err := assetsTraversal(start)
		if err != nil {
			panic(err)
		}
	}
}

func assetsTraversal(startDIR string) error {
	assets, err := os.ReadDir(startDIR)
	if err != nil {
		return err
	}
	for _, file := range assets {
		fuleFilePath := filepath.Join(startDIR, file.Name())
		if filepath.Ext(file.Name()) == ".go" {
			GoFile = append(GoFile, fuleFilePath)
		} else if file.IsDir() {
			err := assetsTraversal(fuleFilePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
