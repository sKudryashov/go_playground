package functions

import (
	"fmt"
	"os"
)

func PrintEnvVars()  {
	for _, env := range os.Environ() {
		fmt.Println("env point is: ", env);
	}
}