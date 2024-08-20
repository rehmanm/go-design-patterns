package main

import "fmt"

type PetTurkey struct {
}

func (p *PetTurkey) gobble() {
	fmt.Println("Gobble pet gobble")
}

func (p *PetTurkey) fly() {
	fmt.Println("I stay at home")
}
