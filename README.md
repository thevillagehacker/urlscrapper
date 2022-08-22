# urlscrapy
A web scrapper to extract the URLs embedded on the Website.

## Installation

```sh
go install -v github.com/thevillagehacker/urlscrapy@latest
```

or
              
```sh
git clone https://github.com/thevillagehacker/urlscrapy.git
cd urlscrapy
go build
```
**urlscrapy** binary will be created and Move the binary to the required folder and add the path to the environment variables.

## Usage
```sh
go run urlscrapy.go -u https://example.com
```

or

```sh
urlscrapy -u https://example.com
```

## Update
- [x] Can now pipe the collected URLs to stdin to other tools to check for status codes and others.

### Example
```sh
urlscrapy -u https://example.com | httpx -<flags>
```
