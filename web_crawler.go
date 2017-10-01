package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func doCrawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	if keeper.getOrAdd(url) {

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		urlChannel <- Document{url: url, body: body}
		// fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			doCrawl(u, depth-1, fetcher)
		}
	} else {
		fmt.Printf("    SKIPPING already processed: %s\n", url)
	}
	return
}

func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	doCrawl(url, depth, fetcher)
	close(urlChannel)
	return
}


type SafeKeeper struct {
	visited   map[string]bool
	mux sync.Mutex
}

func (k SafeKeeper) getOrAdd(u string) bool {
	// lock mutex
	k.mux.Lock()
	rv := true
	if _, ok := k.visited[u]; ok {
		rv = false
	} 
	k.visited[u] = true
	// unlock mutex
	k.mux.Unlock()
	return rv
}

type Document struct {
	url string
	body string
}

var keeper = SafeKeeper{visited: make(map[string]bool)}
var urlChannel = make(chan Document)

func main() {
	go Crawl("http://golang.org/", 4, fetcher)
	// fmt.Printf("%v\n", cap(urlChannel))

	for d := range urlChannel {
		fmt.Printf("URL: %v\n", d.url)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
