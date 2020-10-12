package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println("Now you have %g problems.\n", math.Sqrt(7))
  fmt.Println(add(42, 13))
}

func add(x int, y int) int {
  return x + y
}

