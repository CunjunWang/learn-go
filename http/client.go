package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

const url = "http://www.imooc.com"

func main() {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone X)")
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
