package main

import "fmt"

type WildTurkey struct{}

func (w *WildTurkey) gobble() {
	fmt.Println("Gobble gobble")
}

func (w *WildTurkey) fly() {
	fmt.Println("I'm flying a short distance")
}
