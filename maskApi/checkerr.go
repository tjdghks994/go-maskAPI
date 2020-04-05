package maskapi

import (
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStaus(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("staus code error : %d, %s", res.StatusCode, res.Status)
	}
}
