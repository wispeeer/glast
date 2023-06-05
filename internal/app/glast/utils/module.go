package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/wispeeer/glast/pkg/utils"
)

func ModuleName() string {
	if !utils.IsExist("go.mod") {
		panic("go.mod not found")
	}
	file, err := os.Open("go.mod")
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
