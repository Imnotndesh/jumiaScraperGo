package Modules

import (
	"encoding/csv"
	"os"
	"sync"
)

func SaveToCSV(filename string, itemData []ProductInfo, wg *sync.WaitGroup) {
	file, err := os.Create(filename)
	Checkerr(err)
	defer func(file *os.File) {
		err := file.Close()
		Checkerr(err)
	}(file)
	csvWriter := csv.NewWriter(file)
	header := []string{"Item Name", "Item Price", "Item Url"}
	err = csvWriter.Write(header)
	Checkerr(err)
	for _, item := range itemData {
		record := []string{item.Name, item.Price, item.Url}
		err := csvWriter.Write(record)
		Checkerr(err)
	}
	csvWriter.Flush()
	wg.Done()
}
