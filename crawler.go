// lets get started

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://example.com"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyBytes))
}