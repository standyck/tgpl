// Echo3_Ex1.2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for ix, arg := range os.Args[1:] {
		s += fmt.Sprintf("%s%d %s",sep,ix,arg)
		sep = ", "
	}
	fmt.Println(s)
}
