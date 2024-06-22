package main

import (
	"jumiaScraper/Modules"
)

func startScrape() {
	var itemName string
	var saveMethod string
	//flag.StringVar(&itemName, "i", "", "Item name to fetch")
	//flag.StringVar(&saveMethod, "s", "", " 'j' To save to json file")
	//flag.Parse()
	itemName = "Samsung a15"
	saveMethod = "c"
	Modules.GetItemFromJumia(itemName, saveMethod)
}

func main() {
	startScrape()
}
