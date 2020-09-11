package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type UserWallet struct {
	ID            string `json:"_id"`
	InWalletToken string `json:"tokenId"`
	StoreID       int    `json:"storeId"`
	UserID        int    `json:"userId"`
}

type UserWalletList []UserWallet

const (
	source = "dataSource.json"
	path   = "./generatedFiles/"
)

func main() {
	//Get file source
	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Read and parse data into Json
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var jsonData UserWalletList
	json.Unmarshal(bytes, &jsonData)

	//Create container folder when necessary
	go func() {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0755)
		}
	}()

	//Write file for each Json object
	for n := 0; n < len(jsonData); n++ {
		fileData, _ := json.MarshalIndent(jsonData[n], "", "")
		_ = ioutil.WriteFile(fmt.Sprintf("%s/file%v.json", path, n+1), fileData, 0644)
	}
}
