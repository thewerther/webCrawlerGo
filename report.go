package main

import (
  "fmt"
  "sort"
)

func printReport(pages map[string]int, baseURL string) {
  seperator := "============================="
  fmt.Println(seperator)
  fmt.Printf("REPORT for %s\n", baseURL)
  fmt.Println(seperator)

  // sort pages by most visited
  keys := make([]string, 0, len(pages))
  for key := range pages {
    keys = append(keys, key)
  }
  sort.Slice(keys, func(i, j int) bool { return pages[keys[i]] > pages[keys[j]] })

  for _, key := range keys {
    fmt.Printf("Found %v internal links to %v\n", pages[key], key)
  }
}
