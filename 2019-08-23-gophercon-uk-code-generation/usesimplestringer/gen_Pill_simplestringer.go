package main

import "fmt"

func (v Pill) String() string {
	switch v {
	case Aspirin:
		return "Aspirin"
	case Ibuprofen:
		return "Ibuprofen"
	case Paracetamol:
		return "Paracetamol"
	case Placebo:
		return "Placebo"
	default:
		panic(fmt.Errorf("unknown Pill value %d", v))
	}
}
