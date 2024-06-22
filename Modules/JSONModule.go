package Modules

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

func SaveToJSON(itemData interface{}, passedItemName string, wg *sync.WaitGroup) {
	filename := passedItemName + ".json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		data := itemData
		jsonData, err := json.Marshal(data)
		Checkerr(err)
		err = os.WriteFile(filename, jsonData, 0644)
		Checkerr(err)
	} else {
		log.Println("File", filename, "already exists.Skipping generation")
	}
	wg.Done()
}
