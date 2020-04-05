package main

import (
	"fmt"
	maskapi "tjdghks994/maskAPI/maskApi"
)

var salesURL string = "https://8oi9s0nnth.apigw.ntruss.com/corona19-masks/v1/sales/json?page="
var storesURL string = "https://8oi9s0nnth.apigw.ntruss.com/corona19-masks/v1/stores/json?page="

func main() {

	totalStore := maskapi.Page(storesURL)
	totalSales := maskapi.Page(salesURL)

	storeChan := make(chan []maskapi.MaskInfo)
	store := []maskapi.MaskInfo{}

	salesChan := make(chan []maskapi.MaskInfo)
	sales := []maskapi.MaskInfo{}

	for i := 1; i <= totalStore; i++ {
		go maskapi.Scraper("store", storesURL, i, storeChan)
	}
	for i := 1; i <= totalStore; i++ {
		storeTemp := <-storeChan
		store = append(store, storeTemp...)
	}

	for i := 1; i <= totalSales; i++ {
		go maskapi.Scraper("sales", salesURL, i, salesChan)
	}
	for i := 1; i <= totalSales; i++ {
		salesTemp := <-salesChan
		sales = append(sales, salesTemp...)
	}

	fmt.Println(len(store), len(sales))

}
