package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/fatih/color" // Importing the color package for colored output
	banner "github.com/thevillagehacker/urlscrapper/modules"
)

func main() {

	// Show banner
	banner.ShowBanner()

	// Define flags
	url := flag.String("u", "default value", "target url")
	output := flag.String("o", "", "output file name")       // Flag for output file
	statusCheck := flag.Bool("sc", false, "check status codes for URLs") // Flag for status code check
	flag.Parse()
	var target string
	target = *url

	// Make HTTP request to the target URL
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

		// Print URLs and optionally check status codes
		for _, link := range urls {
			if *statusCheck {
				// Send HTTP request to the URL and get the status code
				resp, err := http.Get(link)
				if err != nil {
					log.Printf("Error checking status code for %s: %v\n", link, err)
					if file != nil {
						fmt.Fprintf(file, "Error checking status code for %s: %v\n", link, err)
					}
				} else {
					statusCode := resp.StatusCode
					coloredStatus := getColorStatus(statusCode)

					// Print to console with colored status code
					fmt.Printf("%s - %s\n", link, coloredStatus)

					if file != nil {
						// Write to file without color
						fmt.Fprintf(file, "%s - %d\n", link, statusCode)
					}
					resp.Body.Close()
				}
			} else {
				// Print URLs to console
				fmt.Println(link)

				if file != nil {
					// Write URLs to file
					fmt.Fprintln(file, link)
				}
			}
		}
	}
}

// getColorStatus returns the status code as a string with the appropriate color
func getColorStatus(statusCode int) string {
	switch {
	case statusCode == 200:
		return color.GreenString("%d", statusCode)
	case statusCode == 301 || statusCode == 302:
		return color.YellowString("%d", statusCode)
	case statusCode == 404:
		return color.BlueString("%d", statusCode)
	case statusCode == 403 || statusCode >= 500:
		return color.RedString("%d", statusCode)
	default:
		return fmt.Sprintf("%d", statusCode)
	}
}
