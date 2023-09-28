package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var CURRENCIES = [...]string{
	"bwp",
	"all",
	"mkd",
	"bmd",
	"iqd",
	"bmd",
	"bnd",
	"gbp",
	"rub",
	"sek",
	"ssp",
	"scr",
	"gyd",
	"ttd",
	"ghs",
	"aed",
	"ggp",
	"rsd",
	"uyu",
	"amd",
	"sar",
	"ars",
	"irr",
	"usd",
	"aud",
	"nok",
	"inr",
	"nzd",
	"bgn",
	"dkk",
	"cad",
	"chf",
	"eur",
}

const BaseUrl string = "https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/"

type CurrencyRates struct {
	name string
	currencies map[string]float64
}

func make_get_request(url string) map[string]interface{} {
	var response string
	var response_json map[string]interface{}

	resp, _ := http.Get(url)

	body, _ := ioutil.ReadAll(resp.Body)
	response = string(body)

	json.Unmarshal([]byte(response), &response_json)

	return response_json
}

func get_currency_rates(currency_key string, ch chan CurrencyRates)  {
	var response_json map[string]float64

	curr_url := fmt.Sprintf("%scurrencies/%s.json", BaseUrl, currency_key)
	curr_json := make_get_request(curr_url)

	b, _ := json.Marshal(curr_json[currency_key])
	json.Unmarshal(b, &response_json)

	filtered_json := make(map[string]float64)

	for _, key := range CURRENCIES {
		val, ok := response_json[key]
		if ok {
			filtered_json[key] = val
		}
	}

	res := CurrencyRates{currency_key, filtered_json}
	ch <- res
}

// Generates a graph with conversion rates
func GenerateConversionGraph() map[string]map[string]float64 {
	exchanges := make(map[string]map[string]float64)

	channel := make(chan(CurrencyRates), len(CURRENCIES))
	
	for _, cur := range CURRENCIES {
		go get_currency_rates(cur, channel)
	}
	
	counter := 0
	for counter < len(CURRENCIES) {
		i := <- channel
		exchanges[i.name] = i.currencies
		counter++
	}

	return exchanges

}

