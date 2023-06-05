package usage

import (
	"fmt"

	"github.com/wispeeer/glast/pkg/info"
)

// app usage menual
func Usage(extraUsage string) {
	fmt.Printf(`Usage: %s -[fhv]

     ------- < Commands Arguments > -------
argument:
  -e, entry          ast analyse entry. 
%s
options:
  -h, help          Show help message. 
  -v, version       Show the app version.
For more help, you can use '%s help' for the detailed information
and you also can check the docs: https://github.com/wispeeer/glast.git  
`, info.MicroService, extraUsage, info.MicroService)
}
