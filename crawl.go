package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

// the buffered chanel models concurrency primitive called counting semaphore
// to limit parallelism
var tokens = make(chan struct{}, 20)

// the function prints the url and return extracted urls
// it starts execution whenever there is a free spot in the chanel
// and frees the spot after it's done
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	links, err := extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return links
}

// the function extracts urls from a page having this url
func extract(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = res.Body.Close()
	}()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", res.StatusCode)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := res.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}
					links = append(links, link.String())
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// utility function that applies function
// passed as parameters to each node in html tree
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
