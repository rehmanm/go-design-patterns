package main

import "fmt"

type MallardDuck struct {
}

func (m *MallardDuck) quack() {
	fmt.Println("Quack")
}
func (m *MallardDuck) fly() {
	fmt.Println("Fly")
}
