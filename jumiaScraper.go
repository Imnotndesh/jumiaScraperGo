package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gocolly/colly/v2"
	"io/ioutil"
	"log"
	"os"
)

// Generic types definition
type productInfo struct {
	Name, Price, Url string
}

var itemData []productInfo

// Saving to JSON logic
func saveToJSON(itemData interface{}, passedItemName string) {
	filename := passedItemName + ".json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("file", filename, "Does not exist. Creating a new one.")
		data := itemData
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling data: ", err)
			return
		}
		err = ioutil.WriteFile(filename, jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file: ", err)
			return
		}
	} else {
		fmt.Println("File", filename, "already exists.Skipping generation")
	}

}

// Scraping logic
func getItemFromJumia(passedItemName, getJson string) {
	url := "http://jumia.co.ke/catalog/?q=" + passedItemName + "&sort=lowest-price&shipped_from=country_local"

	jumiaColly := colly.NewCollector()
	jumiaColly.OnHTML("a.core", func(h *colly.HTMLElement) {

		// Taking product info off page
		itemName := h.ChildText("h3")
		itemPrice := h.ChildText(".prc")
		itemUrl := "Http://jumia" + h.Attr("href")

		collData := productInfo{
			Name:  itemName,
			Price: itemPrice,
			Url:   itemUrl,
		}
		itemData = append(itemData, collData)

	})
	jumiaColly.OnError(func(r *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	jumiaColly.Visit(url)

	if getJson == "y" || getJson == "Y" {
		saveToJSON(itemData, passedItemName)
		fmt.Println("Works")
	} else {
		fmt.Println(itemData)
	}
}

// Flag parsing here
func startScrape() {
	var passedItemName string
	var getJson string

	// Flags are as follows:
	// -i "item name"
	// -s "y" {to save to JSON}

	flag.StringVar(&passedItemName, "i", "", "Item name to fetch")
	flag.StringVar(&getJson, "s", "", " 'y' To save to json file")
	flag.Parse()
	getItemFromJumia(passedItemName, getJson)
}
func main() {
	fmt.Println("...... Fetching ......")
	startScrape()
	println("...... Closing fetcher ......")
}
