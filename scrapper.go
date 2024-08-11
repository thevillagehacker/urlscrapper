package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	banner "github.com/thevillagehacker/urlscrapper/modules"
)

func main() {

	// Show banner
	banner.ShowBanner()

	// Define flags
	url := flag.String("u", "default value", "target url")
	output := flag.String("o", "", "output file name") // New flag for output file
	flag.Parse()
	var target string
	target = *url

	// Make HTTP request
	response, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read response data into memory
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
		var file *os.File
		if *output != "" {
			// Open file for writing
			file, err = os.Create(*output)
			if err != nil {
				log.Fatal("Error creating output file. ", err)
			}
			defer file.Close()
		}

		// Print URLs to both console and file (if specified)
		for _, links := range urls {
			fmt.Println(links) // Print to console
			if file != nil {
				fmt.Fprintln(file, links) // Write to file
			}
		}
	}
}
