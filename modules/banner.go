package banner

import (
	"fmt"
)

// banner
const banner = `
            	URL Scrapper
             ------------------
          ~ |Do Hacks to Secure| ~
             ------------------             v1.1
                    By
`

// Version is the current version of urlscrapy
const Version = `v1.1`

// showBanner is used to show the banner to the user
func ShowBanner() {
	fmt.Printf("%s\n", banner)
	fmt.Printf("\tThe Village Hacker Security\n\n")

	fmt.Printf("Use with caution. You are responsible for your actions.\n")
	fmt.Printf("Developers assume no liability and are not responsible for any misuse or damages.\n")
	fmt.Printf("--------------------------------------------------------------------------------- \n")
}
