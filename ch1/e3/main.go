// John Meuser's Solution to Exercise 1.3 of TGPL
// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var echos = [3]func(){echo1,echo2,echo3}

func main() {
	for i,f:= range echos {
		fmt.Printf("echo%d: %s\n",1+i,stopwatch(f).String())
	}
}

func stopwatch(f func()) time.Duration {
	t0 := time.Now()
	f()
	return time.Since(t0)
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
