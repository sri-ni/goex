package main

import (
  "fmt"
)

func printer(aa []string) {
  fmt.Printf("%d\n", len(aa))
  for i, val := range aa {
    fmt.Printf("%d %s\n", i, val)
  }
}

func main() {
  fmt.Printf("Hello Worlds\n")

  words := make([]string, 0, 5)
  words = append(words, "srini")
  words = append(words, "riddhi")
  printer(words)

  newWords := make([]string, 0, 5)
  copy(newWords, words)
  newWords = append(newWords, "rishi")
  // newWords[1] = "rishi"
  printer(newWords)
}
