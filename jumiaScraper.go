package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)
type productInfo struct{
	Name ,Price, url string
}
func main(){
	println("Starting task")
	jumiaColly := colly.NewCollector()
	url := "http://jumia.co.ke/catalog/?q=shoes&sort=lowest-price&shipped_from=country_local"
	
	jumiaColly.OnHTML("a.core", func(h *colly.HTMLElement) {
		// Instance of the struct above
		prodInfo := productInfo{}

		// Taking product info off page
		prodInfo.Name = h.ChildText("h3")
		prodInfo.Price = h.ChildText(".prc")
		prodInfo.url = "Http://jumia"+h.Attr("href")

		fmt.Printf("Name: %s, \nPrice: %s , \nUrl: %s\n",prodInfo.Name,prodInfo.Price,prodInfo.url )
		fmt.Println("....")
	})
	jumiaColly.OnError(func(r *colly.Response, err error) {
		log.Println("Error: ",err)
	})

	err := jumiaColly.Visit(url)
	if err != nil {
		log.Fatal("Error visiting url: ",err)
	}
	println("Ended task")
}
