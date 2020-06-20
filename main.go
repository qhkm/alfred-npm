package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Resp ...
type Resp struct {
	Items []Item `json:items`
}

// Item ...
type Item struct {
	ID   uint   `json:id`
	Name string `json:name`
}

func fetchData(query string) (*http.Response, error) {
	var npmURL string

	npmURL = "https://api.github.com/search/repositories" + "?q=" + query + "&sort=stars&order=desc"

	// make API request
	resp, err := http.Get(npmURL)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func main() {

	var queryArg string

	if len(os.Args) > 1 {
		queryArg = os.Args[1]
	} else {
		log.Fatalln("no arg")
		return
	}

	resp, err := fetchData(queryArg)
	if err != nil {
		log.Fatalln()
	}

	defer resp.Body.Close()

	// read body into buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var obj Resp

	if err := json.Unmarshal(body, &obj); err != nil {
		panic(err)
	}
	for _, item := range obj.Items {
		fmt.Println(item.ID)
	}

}
