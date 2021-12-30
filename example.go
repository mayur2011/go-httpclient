package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mayur2011/go-httpclient/gohttp"
)

//Singleton http client
var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
}

func getUrls() {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")

	response, err := githubHttpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
