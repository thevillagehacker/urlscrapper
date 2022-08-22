package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	banner "github.com/thevillagehacker/urlscrapy/modules"
)

func main() {

	//banner
	banner.ShowBanner()

	//flags
	url := flag.String("u", "default value", "target url")
	flag.Parse()
	var target string
	target = *url

	// Make HTTP request
	response, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	// Create a regular expression to find URLs
	re := regexp.MustCompile("(http|https)://[a-zA-Z0-9./?=_=-]*")
	urls := re.FindAllString(string(body), -1)
	if urls == nil {
		fmt.Println("No matches.")
	} else {
		for _, links := range urls {
			fmt.Println(links)
		}
	}
}
