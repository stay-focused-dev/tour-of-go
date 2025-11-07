package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Cache struct {
	urls map[string]bool
	mu   sync.Mutex
}

func (c *Cache) Insert(u string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.urls[u]; ok {
		return false
	}

	c.urls[u] = true
	return true
}

func Crawl(url string, depth int, fetcher Fetcher) {
	cache := &Cache{urls: make(map[string]bool)}
	var wg sync.WaitGroup
	crawl(url, depth, fetcher, cache, &wg)
	wg.Wait()
}

func crawl(url string, depth int, fetcher Fetcher, cache *Cache, wg *sync.WaitGroup) {
	if depth <= 0 {
		return
	}

	if !cache.Insert(url) {
		fmt.Printf("Url %v found in cache\n", url)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			crawl(u, depth-1, fetcher, cache, wg)
		}
	}()
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

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
