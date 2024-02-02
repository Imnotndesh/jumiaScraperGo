package main

import (
	"fmt"
	"log"
	"github.com/gocolly/colly/v2"
)
type productInfo struct{
	Name ,Price, url string
}
func fetcher(itemName string){
	jumiaColly := colly.NewCollector()
	url := "http://jumia.co.ke/catalog/?q="+itemName+"&sort=lowest-price&shipped_from=country_local"
	
	jumiaColly.OnHTML("a.core", func(h *colly.HTMLElement) {
		// Instance of the struct above
		prodInfo := productInfo{}

		// Taking product info off page
		prodInfo.Name = h.ChildText("h3")
		prodInfo.Price = h.ChildText(".prc")
		prodInfo.url = "Http://jumia"+h.Attr("href")
		fmt.Println(" ")
		fmt.Println("....")
		fmt.Printf("Name: %s, \nPrice: %s , \nUrl: %s\n",prodInfo.Name,prodInfo.Price,prodInfo.url )
		fmt.Println("....")
		fmt.Println(" ")
	})
	jumiaColly.OnError(func(r *colly.Response, err error) {
		log.Println("Error: ",err)
	})
	jumiaColly.Visit(url)
}
func main(){
	println("Starting fetcher...")
	var passedUserInput string
	println("Enter an item name or category... ")
	fmt.Scanln(&passedUserInput)
	fmt.Println("Finding..."+passedUserInput)
	fetcher(passedUserInput)
	println(" ")
	println("Closing fetcher")
}
