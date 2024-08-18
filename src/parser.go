package main

import (
    "golang.org/x/net/html"
    "log"
    "strings"
)

// ParseLinks extracts all HTTP/HTTPS links from the HTML content.
func ParseLinks(url string, htmlContent string) []string {
    doc, err := html.Parse(strings.NewReader(htmlContent))
    if err != nil {
        log.Fatalf("Error parsing HTML: %v", err)
    }

    var links []string
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    link := attr.Val
                    if strings.HasPrefix(link, "http") {
                        links = append(links, link)
                    }
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
    return links
}
