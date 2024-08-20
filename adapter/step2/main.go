package main

import "fmt"

func main() {

	duck := &MallardDuck{}
	wildTurkey := &WildTurkey{}
	wildTurkeyAdapter := &TurkeyAdapter{
		turkey: wildTurkey,
	}

	petTurkey := &PetTurkey{}
	petTurkeyAdapter := &TurkeyAdapter{
		turkey: petTurkey,
	}

	fmt.Println("The Wild Turkey says...")
	wildTurkey.gobble()
	wildTurkey.fly()

	fmt.Println("\nThe Duck says...")
	testDuck(duck)

	fmt.Println("\nThe TurkeyAdapter says...")
	testDuck(wildTurkeyAdapter)

	fmt.Println("\nThe PetTurkeyAdapter says...")

	testDuck(petTurkeyAdapter)

}

func testDuck(duck Duck) {
	duck.quack()
	duck.fly()
}
