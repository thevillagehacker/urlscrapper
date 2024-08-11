# urlscrapper
A web scrapper to extract the URLs embedded on the Website.

## Installation

```sh
go install -v github.com/thevillagehacker/urlscrapper@latest
```

or
              
```sh
git clone https://github.com/thevillagehacker/urlscrapper.git
cd urlscrapper
go build
```

**urlscrapper** binary will be created and Move the binary to the required folder and add the path to the environment variables.

## Usage
```sh
go run scrapper.go -u https://example.com
```

or

```sh
urlscrapper -u https://example.com
```

## Update
- [x] Can now pipe the collected URLs to stdin to other tools to check for status codes and others.
- [X] Can now output the scrapped data into a file using `-o` flag.

### Example
```sh
urlscrapper -u https://example.com | httpx -<flags>
```

```sh
urlscrapper -u https://example.com -o out.txt
```