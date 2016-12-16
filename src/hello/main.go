package main

import (
  "fmt";
  "os"
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

  mapMaker()

  byteSlicer()
}

func mapMaker() {
  monthr := make(map[string]int)
  monthr["jan"] = 31
  monthr["feb"] = 28
  monthr["mar"] = 31
  for mo, da := range monthr {
    fmt.Printf("%s has %d days\n", mo, da)
  }

  ftale := map[string]string {
    "jack": "jill",
    "fred": "wilma",
    "hansel": "gretel",
  }
  for na, da := range ftale {
    fmt.Printf("%s goes with %s\n", na, da)
  }
}

func byteSlicer() {
  b := make([]byte, 100)

  f, err := os.Open("test.txt")
  f.Read(b)
  str := string(b[:])
  if err != nil {
    fmt.Printf("Read string = %d %s\n", len(b), str)
  }
}
