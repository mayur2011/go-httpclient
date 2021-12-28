package main

import "github.com/mayur2011/go-httpclient/gohttp"

func basicExample() {

	client := gohttp.New()
	client.Get()
}
