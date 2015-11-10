// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	// BobK:  This assignment looks funky and error prone
	s, sep := "", ""
	// BobK:  It is interesting to note that range is returning a tuple
	// and that we are ignoring, through the zot, the first element
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ", "
	}
	fmt.Println(s)
}
