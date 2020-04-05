package maskapi

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

//Scraper scraper info
func Scraper(t, url string, num int, c chan []MaskInfo) {
	res, err := http.Get(url + strconv.Itoa(num))
	checkErr(err)
	checkStaus(res)
	defer res.Body.Close()

	m := []MaskInfo{}

	switch t {
	case "sales":
		mask := maskSales{}
		dec := json.NewDecoder(res.Body)
		dec.Decode(&mask)
		for i := 0; i < mask.Cnt; i++ {
			temp := MaskInfo{
				Code:   mask.Sale[i].Code.String(),
				Remain: mask.Sale[i].Remain,
			}
			m = append(m, temp)

		}

	case "store":
		mask := maskStore{}
		dec := json.NewDecoder(res.Body)
		dec.Decode(&mask)
		for i := 0; i < mask.Cnt; i++ {
			temp := MaskInfo{
				Code: mask.Store[i].Code.String(),
				Name: mask.Store[i].Name,
				Addr: mask.Store[i].Addr,
				Lat:  mask.Store[i].Lat.String(),
				Lng:  mask.Store[i].Lng.String(),
			}
			m = append(m, temp)
		}
	}
	c <- m
}

func writeCSV(m []MaskInfo) {

	file, err := os.Create("maskInfo.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"Code", "Name", "Remain", "Addr", "Lat", "Lng"}
	wErr := w.Write(header)
	checkErr(wErr)

	for _, mask := range m {
		maskStr := []string{mask.Code, mask.Name, mask.Remain, mask.Addr, mask.Lat, mask.Lng}
		mErr := w.Write(maskStr)
		checkErr(mErr)
	}
}

//Page total pages
func Page(url string) int {
	res, err := http.Get(url + "1")
	checkErr(err)
	checkStaus(res)
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	checkErr(readErr)

	m := maskStore{}
	jsonErr := json.Unmarshal(body, &m)
	checkErr(jsonErr)

	return m.TotalPages
}
