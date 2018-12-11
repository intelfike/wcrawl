package main

import (
	// "crypto/tls"
	// "net/http"
	// "fmt"
	// "strings"

	"github.com/intelfike/wcrawl"
)

func main() {
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// links, _ := wcrawl.GetLinks("https://localplace.jp/")
	// fmt.Println(strings.Join(links, "\n"))
	wc := wcrawl.Crawler{}
	wc.Do("https://localplace.jp/")
}
