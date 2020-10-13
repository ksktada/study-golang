package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println("Now you have %g problems.\n", math.Sqrt(7))
  fmt.Println(add(42, 13))
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}

// func
func add(x int, y int) int {
  return x + y
}

// multiple results
func swap(x, y string) (string, string) {
  return  y, x
}
