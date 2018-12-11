package wcrawl

import (
	"fmt"
	// "io/ioutil"
	// "net/http"
	"regexp"
	"os/exec"
	"strings"
	"time"
)

// type Site struct {
// 	URL string
// 	Title string
// }
var urlReg = regexp.MustCompile(`href=["'][^"']+["']`)

type Crawler struct {
	URLs map[string]string
	Site string
}

func (c *Crawler) Do(site string) map[string]string {
	c.Site = site
	c.URLs = map[string]string{}
	c.URLs[site] = ""
	c.recCrawl(site)

	return c.URLs
}

func (c *Crawler) recCrawl(url string) {
	urls := c.GetLinks(url)
	_ = urls
	for link, _ := range urls {
		time.Sleep(time.Second / 10)
		fmt.Println(link)

		c.recCrawl(link)
	}
}

// 未探索のリンクのみを返す
func (c *Crawler) GetLinks(url string) map[string]string {
	links, _ := GetLinks(url)
	newlink := map[string]string{}
	for _, link := range links {
		if link == "" {
			continue
		}
		if strings.HasPrefix(link, "/") {
			link = strings.TrimSuffix(c.Site, "/") + link
		}

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
