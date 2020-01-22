package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	fGreeting = flag.String("greeting", "Hello", "The greeting to use")
	fTarget   = flag.String("target", "World", "Who we are greeting")
)

func main() {
	os.Exit(main1())
}

func main1() int {
	if err := mainerr(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func mainerr() error {
	flag.Parse()
	fmt.Printf("%v, %v!\n", *fGreeting, *fTarget)
	return nil
}
