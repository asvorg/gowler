package main

import (
    "flag"
    "log"
    "sync"
)

func main() {
    // Define command-line flags
    var startUrl string
    var depth int
    flag.StringVar(&startUrl, "url", "http://example.com", "Starting URL for the crawler")
    flag.IntVar(&depth, "depth", 1, "Depth of crawling")
    flag.Parse()

    // Validate depth
    if depth < 1 {
        log.Fatalf("Depth must be a positive integer")
    }

    // Initialize visited map and mutex
    visited := make(map[string]bool)
    var mu sync.Mutex

    // Start crawling
    Crawl(startUrl, depth, visited, &mu)
}
