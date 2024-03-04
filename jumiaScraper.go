package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
)

type productInfo struct {
	Name, Price, url string
}

var itemData []productInfo
//var allData []productInfo

func getItemFromJumia(url string) {
	jumiaColly := colly.NewCollector()

	jumiaColly.OnHTML("a.core", func(h *colly.HTMLElement) {

		// Taking product info off page
		itemName := h.ChildText("h3")
		itemPrice := h.ChildText(".prc")
		itemUrl := "Http://jumia" + h.Attr("href")

		collData := productInfo{
			Name: itemName,
			Price: itemPrice,
			url: itemUrl, 
		}
		itemData = append(itemData, collData)

	})
	jumiaColly.OnError(func(r *colly.Response, err error) {
		log.Println("Error: ", err)
	})
	jumiaColly.Visit(url)
	fmt.Println(itemData[0])
}
func generateUrl(itemName string) {
	url := "http://jumia.co.ke/catalog/?q=" + itemName + "&sort=lowest-price&shipped_from=country_local"
	getItemFromJumia(url)
}

func startScrape() {
	var passedUserInput string
	fmt.Println("Fetching.......")
	flag.StringVar(&passedUserInput, "item", "", "Item to fetch")
	flag.Parse()
	generateUrl(passedUserInput)
}
func main() {
	startScrape()
	println("Closing fetcher")
}
