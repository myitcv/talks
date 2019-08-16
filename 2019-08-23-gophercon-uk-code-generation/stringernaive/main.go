package main

import "fmt"

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
)

// stringfn OMIT
func (p Pill) String() string {
	switch p {
	case Aspirin:
		return "Aspirin"
	case Ibuprofen:
		return "Ibuprofen"
	case Paracetamol:
		return "Paracetamol"
	case Placebo:
		return "Placebo"
	}
	panic("oh dear")
}

// stringfnend OMIT

func main() {
	var p Pill = Ibuprofen
	fmt.Printf("You need to take 2 %v per day\n", p)
}
