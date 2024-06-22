package Modules

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"sync"
)

type ProductInfo struct {
	Name, Price, Url string
}

var ItemData []ProductInfo

func Checkerr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func logErr(err error) {
	log.Println("Error: ", err)
}

func GetItemFromJumia(itemName, saveMethod string) {
	url := "http://jumia.co.ke/catalog/?q=" + itemName + "&sort=lowest-price&shipped_from=country_local"
	jumiaColly := colly.NewCollector()
	jumiaColly.OnHTML("a.core", func(h *colly.HTMLElement) {
		itemName := h.ChildText("h3")
		itemPrice := h.ChildText(".prc")
		itemUrl := "Http://jumia" + h.Attr("href")

		collData := ProductInfo{
			Name:  itemName,
			Price: itemPrice,
			Url:   itemUrl,
		}
		ItemData = append(ItemData, collData)
	})

	jumiaColly.OnError(func(r *colly.Response, err error) {
		logErr(err)
	})
	err := jumiaColly.Visit(url)
	Checkerr(err)
	var wg sync.WaitGroup
	wg.Add(1)
	if saveMethod == "j" || saveMethod == "J" {
		go SaveToJSON(ItemData, itemName, &wg)
	} else if saveMethod == "c" {
		filename := itemName + ".csv"
		go SaveToCSV(filename, ItemData, &wg)
	} else {
		for i := 0; i < len(ItemData); i++ {
			fmt.Println(ItemData[i])
		}
	}
	wg.Wait()
}
