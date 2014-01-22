package magic

import (
  "log"
  "net/http"
  "strings"

  "code.google.com/p/go.net/html"
  "code.google.com/p/go.net/html/atom"
)

func punctuation(c rune) rune {
  switch c {
  case '.', '?', ',', '-', '$', '(', ')',
    '{', '}', '[', ']', '\'', ';', '|',
    '&', '“', '’', '/', '=', '—', '"':
    return -1
  default:
    return c
  }
}

func findText(n *html.Node, words []string) []string {
  if n.Parent != nil && n.Parent.DataAtom == atom.P && n.Type == html.TextNode {
    for _, word := range strings.Split(n.Data, " ") {
      words = append(words, word)
    }
  }
  if n.FirstChild != nil {
    words = findText(n.FirstChild, words)
  }
  if n.NextSibling != nil {
    words = findText(n.NextSibling, words)
  }
  return words
}

func getText(url string) []string {
  resp, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  doc, err := html.Parse(resp.Body)
  if err != nil {
    log.Fatal(err)
  }
  return findText(doc, make([]string, 0))
}

func GetWords(urls []string) []string {
  words := make(chan string)
  go func() {
    for _, url := range urls {
      for _, word := range getText(url) {
        word := strings.TrimSpace(strings.ToLower(strings.Map(punctuation, word)))
        if word != "" {
          words <- word
        }
      }
    }
    close(words)
  }()
  s := make([]string, 0)
  for word := range words {
    s = append(s, word)
  }
  return s
}
