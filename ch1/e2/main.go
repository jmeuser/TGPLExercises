// John Meuser's Solution to Exercise 1.2 of TGPL
// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, x := range os.Args {
		fmt.Println(i,x)
	}
}
