package main

import "fmt"

//go:generate simplestringer Pill

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
)

func main() {
	var p Pill = Ibuprofen
	fmt.Printf("You need to take 2 %v per day\n", p)
}
