// Echo3_Ex1.1 prints its command-line arguments with Args[0]
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
