package main

import (
	"fmt"

	"github.com/mayur2011/go-httpclient/gohttp"
)

func main() {

	client := gohttp.New()
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
}
