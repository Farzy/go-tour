package main

import (
	"fmt"
	"go-tour/utils"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slide of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

type SafeUrl struct {
	urls map[string]bool
	mux  sync.Mutex
}

// Check if an URL has already been visited, add it to the list
// in any case.
func (u *SafeUrl) visit(url string) bool {
	u.mux.Lock()
	defer u.mux.Unlock()

	_, ok := u.urls[url]
	if ok {
		// URL already visited
		return false
	} else {
		// Add URL to list and visit it
		u.urls[url] = true
		return true
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum depth.
func Crawl(url string, depth int, fetcher Fetcher, urlsVisited *SafeUrl) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if urlsVisited.visit(u) {
			go Crawl(u, depth-1, fetcher, urlsVisited)
		}
	}
	return
}

func exerciceWebCrawler() {
	var DEBUG = utils.Contains(
		[]string{"true", "yes", "y", "1"},
		strings.ToLower(os.Getenv("DEBUG")))
	var DEPTH int
	i, ok := strconv.ParseInt(os.Getenv("DEPTH"), 10, 0)
	if ok != nil || i < 1 {
		DEPTH = 4
	} else {
		DEPTH = int(i)
	}

	fmt.Printf(
		"You can set DEPTH (current: %v) or DEBUG (current: %v) as environment variables\n",
		DEPTH, DEBUG)

	urlsVisited := SafeUrl{
		urls: make(map[string]bool),
	}

	initialProcs := runtime.NumGoroutine()
	urlsVisited.visit("https://golang.org/") // Mark base url as visited
	Crawl("https://golang.org/", DEPTH, fetcher, &urlsVisited)
	for procs := runtime.NumGoroutine(); procs != initialProcs; {
		if DEBUG {
			fmt.Println("Current number of Goroutines:", procs)
		}
		procs = runtime.NumGoroutine()
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
