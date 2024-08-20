package main

type TurkeyAdapter struct {
	turkey Turkey
}

func (t *TurkeyAdapter) quack() {
	t.turkey.gobble()
}

func (t *TurkeyAdapter) fly() {
	for i := 0; i < 5; i++ {
		t.turkey.fly()
	}
}
