package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
  "strings"
)

func getHTML(rawURL string) (string, error) {
  resp, err := http.Get(rawURL)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
    return "", fmt.Errorf("GET failed with status code: %v", resp.StatusCode)
  }

  if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
    return "", errors.New("content type not 'text/html'!\n")
  }

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  return string(body), nil
}
