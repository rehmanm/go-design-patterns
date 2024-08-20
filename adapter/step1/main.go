package main

func main() {

	client := &Client{}
	//mac := &Mac{}
	windowsMachine := &Windows{}

	//client.insertLightningConnectorIntoComputer(mac)
	client.insertLightningConnectorIntoComputer(&WindowsAdapter{windowsMachine})
}
