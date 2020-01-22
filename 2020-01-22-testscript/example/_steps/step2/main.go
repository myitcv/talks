package main

import (
	"flag"
	"fmt"
)

var (
	fGreeting = flag.String("greeting", "Hello", "The greeting to use")
	fTarget   = flag.String("target", "world", "Who we are greeting")
)

func main() {
	flag.Parse()
	fmt.Printf("%v, %v!\n", *fGreeting, *fTarget)
}
