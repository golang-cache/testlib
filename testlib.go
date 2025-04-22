package testlib

import (
	"fmt"
	"net/http"
)

var URL = "https://example.com"
var URL2 = "https://other-url.com"

func PerformGetRequest() {
	resp, err := http.Get(URL)
	http.Get(URL2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Status:", resp.Status)
}
