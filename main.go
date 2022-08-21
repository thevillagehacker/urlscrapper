package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/projectdiscovery/gologger"
)

func main() {

	//banner
	showBanner()

	//flags
	url := flag.String("u", "default value", "a string for description")
	flag.Parse()

	//URL
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

const banner = `
                  URL Scrapy
             ------------------
          ~ |Do Hacks to Secure| ~
             ------------------             v1.0
                    By
`

// Version is the current version of httpx
const Version = `v1.0`

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\tThe Village Hacker Security\n\n")

	gologger.Print().Msgf("Use with caution. You are responsible for your actions.\n")
	gologger.Print().Msgf("Developers assume no liability and are not responsible for any misuse or damages.\n")
}
