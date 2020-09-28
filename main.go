// count tallies the number of times each line
// occurs within a file.
package main

import (
  "fmt"
  "github.com/nparsons/datafile"
  "log"
)

func main() {
  lines, err := datafile.GetStrings("votes.txt")
  if err != nil {
    log.Fatal(err)
  }

  votes := make(map[string]int)
  for _, name := range lines {
    votes[name]++
  }

  for name, count := range votes {
    fmt.Printf("Votes for %s: %d\n", name, count)
  }

}
