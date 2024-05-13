package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/gocolly/colly/v2"
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

	if getJson == "j" || getJson == "J" {
		saveToJSON(itemData, passedItemName)
		fmt.Println("Works")
	}else if getJson == "c" {
		filename := passedItemName+".csv"
		file,err := os.Create(filename)
		if err != nil{
			log.Fatal(err)
		}
		defer file.Close()
		csvWriter := csv.NewWriter(file)
		header := []string{"Item Name","Item Price","Item Url"}
		err = csvWriter.Write(header)
		if err != nil{
			log.Fatal(err)
		}
		for _, item := range itemData{
			record := []string{item.Name,item.Price,item.Url}
			err := csvWriter.Write(record)
			if err != nil{
				log.Fatal(err)
			}
		}
		csvWriter.Flush()
	} else{
		for i :=0; i<len(itemData);i++{
			fmt.Println(itemData[i])
		}
	}
}

// Flag parsing here
func startScrape() {
	var passedItemName string
	var getJson string

	// Flags are as follows:
	// -i "item name"
	// -s "j" {to save to JSON}

	flag.StringVar(&passedItemName, "i", "", "Item name to fetch")
	flag.StringVar(&getJson, "s", "", " 'j' To save to json file")
	flag.Parse()
	getItemFromJumia(passedItemName, getJson)
}

func main() {
	startScrape()
}
