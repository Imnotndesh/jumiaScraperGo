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

func getItemFromJumia(url string) {
	jumiaColly := colly.NewCollector()
	jumiaColly.OnHTML("a.core", func(h *colly.HTMLElement) {
		// Instance of the struct above
		prodInfo := productInfo{}

		// Taking product info off page
		prodInfo.Name = h.ChildText("h3")
		prodInfo.Price = h.ChildText(".prc")
		prodInfo.url = "Http://jumia" + h.Attr("href")
		fmt.Println(" ")
		fmt.Println("....")
		fmt.Printf("Name: %s, \nPrice: %s , \nUrl: %s\n", prodInfo.Name, prodInfo.Price, prodInfo.url)
		fmt.Println("....")
	})
	jumiaColly.OnError(func(r *colly.Response, err error) {
		log.Println("Error: ", err)
	})
	jumiaColly.Visit(url)
}
func generateUrl(itemName string) {
	url := "http://jumia.co.ke/catalog/?q=" + itemName + "&sort=lowest-price&shipped_from=country_local"
	getItemFromJumia(url)
}

func startScrape() {
	var passedUserInput string
	fmt.Println("Fetching.......")
	flag.StringVar(&passedUserInput, "item", "","Item to fetch")
	flag.Parse()
	generateUrl(passedUserInput)
}
func main() {
	startScrape()
	println("Closing fetcher")
}
