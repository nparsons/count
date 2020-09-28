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

  names := make(map[string]int)
  for _, name := range lines {
    names[name]++
  }

  fmt.Println(names)

}
