package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var considered_currencies = [...]string{
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
	"fim",
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
	"ved",
	"bgn",
	"dkk",
	"cad",
	"chf",
	"eur",
}

const base_url string = "https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/"

func make_get_request(url string) map[string]interface{} {
	var response string
	var response_json map[string]interface{}

	resp, err := http.Get(url)

	if err != nil {
		// todo: implement error handling
	}

	body, _ := ioutil.ReadAll(resp.Body)
	response = string(body)

	json.Unmarshal([]byte(response), &response_json)

	return response_json
}

func get_all_currency_values(base_url string) []string {
	var currencies_url = fmt.Sprintf("%scurrencies.json", base_url)
	curr := make_get_request(currencies_url)

	var currency_keys []string

	for key := range curr {
		currency_keys = append(currency_keys, key)
	}

	return currency_keys
}

func get_currency_rates(currency_key string, base_url string) map[string]float32 {
	var response_json map[string]float32

	curr_url := fmt.Sprintf("%scurrencies/%s.json", base_url, currency_key)
	curr_json := make_get_request(curr_url)

	b, _ := json.Marshal(curr_json[currency_key])
	json.Unmarshal(b, &response_json)

	return response_json
}

func main() {
	// currencies := get_all_currency_values(base_url)

	exchanges := make(map[string]map[string]float32)
	for _, currency := range considered_currencies {
		currency_exchange := get_currency_rates(currency, base_url)
		exchanges[currency] = currency_exchange
	}
	fmt.Println(exchanges)
}
