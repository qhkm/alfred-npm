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

func main() {

	var queryArg string

	if len(os.Args) > 1 {
		queryArg = os.Args[1]
	} else {
		log.Fatalln("no arg")
		return
	}
	var npmURL string

	npmURL = "https://api.github.com/search/repositories" + "?q=" + queryArg + "&sort=stars&order=desc"

	// make API request
	resp, err := http.Get(npmURL)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// var prettyJSON bytes.Buffer

	// read body into buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// error := json.Indent(&prettyJSON, body, "", "\t")
	// if error != nil {
	// 	log.Println("JSON parse error")
	// }

	//log.Println(prettyJSON.Bytes())
	var obj Resp

	if err := json.Unmarshal(body, &obj); err != nil {
		panic(err)
	}

	// fmt.Printf("%v\n", obj)

	for _, item := range obj.Items {
		fmt.Println(item.ID)
	}

}
