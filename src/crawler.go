package main

import (
    "fmt"
    "sync"
)

// Crawl recursively crawls the web pages starting from the given URL up to a specified depth.
func Crawl(url string, depth int, visited map[string]bool, mu *sync.Mutex) {
    if depth <= 0 {
        return
    }

    mu.Lock()
    if visited[url] {
        mu.Unlock()
        return
    }
    visited[url] = true
    mu.Unlock()

    fmt.Println("Crawling:", url)
    content := Fetch(url)
    links := ParseLinks(url, content)

    var wg sync.WaitGroup

    for _, link := range links {
        wg.Add(1)
        go func(link string) {
            defer wg.Done()
            Crawl(link, depth-1, visited, mu)
        }(link)
    }

    wg.Wait()
}
