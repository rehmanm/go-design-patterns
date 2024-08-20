package main

import "fmt"

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}
