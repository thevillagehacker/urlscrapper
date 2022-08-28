package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"os"
	"log"
	"sync"

	banner "github.com/thevillagehacker/urlscrapy/modules"
)


func main() {

	//banner
	banner.ShowBanner()
	var domains []string

	//flags
	furl := flag.String("u","", "target url")
	filePtr := flag.String("l","","list of URLs in file")

	flag.Parse()

    f, err := os.Open(*filePtr)
    if err != nil {
        // log.Fatal(err)
    }
    defer func() {
        if err = f.Close(); err != nil {
        // log.Fatal(err)
    }
    }()
    s := bufio.NewScanner(f)
    for s.Scan() {
        domains = append(domains, s.Text())
    }
    err = s.Err()
    if err != nil {
        // log.Fatal(err)
    }
  
	var wg sync.WaitGroup

	for _, u := range domains {
	  
	  wg.Add(1)
	  go func(url string) {
	  
		defer wg.Done()

		response, err := http.Get(url)
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

	  }(u)
	}
	wg.Wait()

	// Make HTTP request
	response, err := http.Get(*furl)
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

