package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/wispeeer/glast/pkg/utils"
)

func ModuleName(entry string) string {
	gomodFilePath := filepath.Join(entry, "go.mod")
	if !utils.IsExist(gomodFilePath) {
		panic("go.mod not found")
	}
	file, err := os.Open(gomodFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var module string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		if strings.HasPrefix(lineText, "module") {
			module = strings.TrimPrefix(lineText, "module")
		}
	}
	module = strClean(module)
	return module
}

func strClean(s string) string {
	str := strings.Join(strings.Split(s, "\""), "")
	str = strings.Join(strings.Fields(str), " ")
	return str
}
