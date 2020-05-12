package alphavantage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

//GetStock Get a stock by symbol
func GetStock(symbol string) (interface{}, error) {
	url := getURL("GLOBAL_QUOTE")
	resp, err := http.Get(url + "&symbol=" + symbol)
	if err != nil {
		return nil, err
	}

	response, err := getBody(resp)
	return response, err
}

//SearchStock by a search term
func SearchStock(keywords string) (interface{}, error) {
	url := getURL("SYMBOL_SEARCH")
	resp, err := http.Get(url + "&keywords=" + keywords)
	if err != nil {
		return nil, err
	}

	response, err := getBody(resp)
	return response, err
}

func getURL(function string) string {
	mainURL := os.Getenv("ALPHA_VANTAGE_URL") + "/query"
	auth := os.Getenv("ALPHA_VANTAGE_API_KEY_NAME") + "=" + os.Getenv("ALPHA_VANTAGE_API_KEY")
	return mainURL + "?" + auth + "&function=" + function
}

func getBody(response *http.Response) (interface{}, error) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(body, &resp)
	if val, ok := resp["Error Message"]; ok {
		return nil, errors.New(val.(string))
	}

	return resp, nil
}
