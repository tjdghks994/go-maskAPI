package maskapi

import "encoding/json"

//MaskInfo total Info
type MaskInfo struct {
	Code   string
	Name   string
	Remain string
	Addr   string
	Lat    string
	Lng    string
}

type maskStore struct {
	Cnt        int      `json:"count"`
	Page       string   `json:"page"`
	Store      []stores `json:"storeInfos"`
	TotalCnt   int      `json:"totalCount"`
	TotalPages int      `json:"totalPages"`
}

type stores struct {
	Addr string      `json:"addr"`
	Code json.Number `json:"code"`
	Lat  json.Number `json:"lat"`
	Lng  json.Number `json:"lng"`
	Name string      `json:"name"`
	Type string      `json:"type"`
}

type maskSales struct {
	Cnt        int     `json:"count"`
	Page       string  `json:"page"`
	Sale       []sales `json:"sales"`
	TotalCnt   int     `json:"totalCount"`
	TotalPages int     `json:"totalPages"`
}

type sales struct {
	Code    json.Number `json:"code"`
	Created string      `json:"created_at"`
	Remain  string      `json:"remain_stat"`
	Stock   string      `json:"stock_at"`
}
