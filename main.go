package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://kalipare.com:80/hello?name=john wick&age=27"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %v\n", urlString)

	fmt.Printf("protocol: %v\n", u.Scheme)
	fmt.Printf("host: %v\n", u.Host)
	fmt.Printf("path: %v\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %v, age: %v", name, age)
}
