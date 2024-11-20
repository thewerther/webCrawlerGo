package main

import (
	"fmt"
	"strings"
	"golang.org/x/net/html"
  "net/url"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	// parse tree root
	reader := strings.NewReader(htmlBody)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var urls []string
	var getURLs func(*html.Node)
	getURLs = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
          urlRef, err := url.Parse(a.Val)
          if err != nil {
            fmt.Printf("couldn't parse href '%v': %v\n", a.Val, err)
            continue
          }

					resolvedURL := baseURL.ResolveReference(urlRef)
					urls = append(urls, resolvedURL.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			getURLs(c)
		}
	}

	getURLs(doc)

	return urls, nil
}
