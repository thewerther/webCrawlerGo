package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := os.Args[1]
  maxConcurrency := 1
  if len(os.Args) == 3 {
    var err error = nil
    maxConcurrency, err = strconv.Atoi(os.Args[2])
    if err != nil {
      log.Fatal(err)
    }
  }

  maxPages := 10
  if len(os.Args) == 4 {
    var err error = nil
    maxPages, err = strconv.Atoi(os.Args[3])
    if err != nil {
      maxPages = 10
    }
  }

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v\n", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

  printReport(cfg.pages, cfg.baseURL.String())
}
