package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

var fetched map[string]bool

// Crawl uses findLinks to recursively crawl
// pages starting with url, to a maximum of depth.

type crawlResult struct {
	url   string
	urls  []string
	err   error
	depth int
}

// TODO: Fetch URLs in parallel.
// create struct to store result of chan type
// use goroutine to fetchLink; if err continue or return err
// mark true in map for url
// use fetching variable as loop & inside get result from ch
// if depth not 0; loop through urls and call crawl
func Crawl(url string, depth int) {
	ch := make(chan *crawlResult)

	fetch := func(url string, depth int) {
		urls, err := findLinks(url)

		ch <- &crawlResult{url, urls, err, depth}
	}

	go fetch(url, depth)
	fetched[url] = true

	for fetching := depth - 1; fetching > 0; fetching-- {
		result := <-ch

		if result.err != nil {
			fmt.Println(result.err)
			continue
		}

		fmt.Printf("found: %s\n", url)

		if result.depth > 0 {
			for _, u := range result.urls {
				if !fetched[u] {
					go fetch(u, result.depth-1)
					fetched[u] = true
					fetching++
				}
			}
		}
	}

	close(ch)
}

func CrawlSequential(url string, depth int) {
	if depth < 0 {
		return
	}
	urls, err := findLinks(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s\n", url)
	fetched[url] = true
	for _, u := range urls {
		if !fetched[u] {
			CrawlSequential(u, depth-1)
		}
	}
}

// Web crawler to process all urls of url according to input depth
func crawler() {
	// fmt.Println("no. of cpus: ", runtime.NumCPU())
	// runtime.GOMAXPROCS(runtime.NumCPU())
	fetched = make(map[string]bool)
	now := time.Now()
	Crawl("https://google.com", 2)
	fmt.Println("time taken:", time.Since(now))

	fetched = make(map[string]bool)
	now = time.Now()
	CrawlSequential("https://google.com", 2)
	fmt.Println("time taken:", time.Since(now))

}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
