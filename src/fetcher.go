package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

// Fetch retrieves the HTML content of a given URL.
func Fetch(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error fetching URL: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %v", err)
    }

    return string(body)
}
