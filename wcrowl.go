package wcrawl

import (
	"fmt"
	// "io/ioutil"
	// "net/http"
	"regexp"
	"os/exec"
	"strings"
)

// type Site struct {
// 	URL string
// 	Title string
// }
var urlReg = regexp.MustCompile(`href=["'][^"']+["']`)

type Crawler struct {
	URLs map[string]string
}

func (c *Crawler) Do(site string) []string {
	c.recCrawl(site)
	fmt.Println(c.URLs)
	return nil
}

func (c *Crawler) recCrawl(url string) []string {
	urls := c.GetLinks(url)
	_ = urls
	return nil // TODO
	for _, url := range urls {
		return c.recCrawl(url)
	}
	return nil
}

// 未探索のリンクのみを返す
func (c *Crawler) GetLinks(url string) map[string]string {
	links, _ := GetLinks(url)

	newlink := map[string]string{}
	for _, link := range links {
		_, ok := c.URLs[link]
		if !ok {
			newlink[link] = url
		}
		c.URLs[link] = url
	}
	return newlink
}

func GetLinks(url string) ([]string, error) {
	cmd := exec.Command("/opt/google/chrome/chrome", "--headless", "--disable-gpu", "--dump-dom", url)
	b, err := cmd.Output()
	// r, err := http.Get(url)
	// if err != nil {
	// 	return nil, err
	// }
	// defer r.Body.Close()
	// b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	ss := urlReg.FindAllString(string(b), -1)
	for n, _ := range ss {
		ss[n] = strings.TrimPrefix(ss[n], `href="`)
		ss[n] = strings.TrimSuffix(ss[n], `"`)
	}

	return ss, nil
}
