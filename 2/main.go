package main

import (
  "fmt"
  "./input"
  "sort"
)

func first(theInput input.Lines) {
  total := 0
  for _, line := range theInput {
    l := line[0]
    w := line[1]
    h := line[2]
    sides := []int {
      l * w, w * h, h * l,
    }

    sort.Ints(sides)

    smallest := sides[0]

    total += (2 * l * w) + (2 * w * h) + (2 * h * l) + smallest
  }

  fmt.Printf("Total: %d\n", total)
}

func second(theInput input.Lines) {
  total := 0
  for _, line := range theInput {
    l := line[0]
    w := line[1]
    h := line[2]
    sides := []int {
      l * w, w * h, h * l,
    }

    sort.Ints(sides)

    smallest := sides[0]
    middle := sides[1]

    total += (smallest * 2) + (middle * 2) + (l * w * h)
  }
  fmt.Printf("Total: %d\n", total)
}

func main() {
  theInput := input.GetInput()
  fmt.Printf("First: \n")
  first(theInput)
  fmt.Printf("Second: \n")
  second(theInput)
}
