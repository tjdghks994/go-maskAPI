package main

import (
	"fmt"

	maskapi "github.com/tjdghks994/go-maskAPI/maskApi"
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

	drugName := "성모약국"
	searchDrug := make(map[string]maskapi.MaskInfo)

	for _, drug := range store {
		if drug.Name == drugName {
			searchDrug[drug.Code] = maskapi.MaskInfo{
				Code: drug.Code,
				Name: drug.Name,
				Addr: drug.Addr,
				Lat:  drug.Lat,
				Lng:  drug.Lng,
			}
		}
	}

	for _, remain := range sales {
		resultKey, resultValue := compareCode(remain, searchDrug)
		searchDrug[resultKey] = resultValue
	}

	for _, v := range searchDrug {
		fmt.Println(v)
	}

}

func compareCode(m maskapi.MaskInfo, drug map[string]maskapi.MaskInfo) (string, maskapi.MaskInfo) {
	for key, value := range drug {
		if m.Code == key {
			temp := maskapi.MaskInfo{}
			temp.Code = key
			temp.Name = value.Name
			temp.Addr = value.Addr
			temp.Lat = value.Lat
			temp.Lng = value.Lng
			temp.Remain = m.Remain

			return key, temp
		}
	}
	return "", maskapi.MaskInfo{}
}
