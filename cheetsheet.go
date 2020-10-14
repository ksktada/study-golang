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
  fmt.Println(split(17))
}

// func
func add(x int, y int) int {
  return x + y
}

// multiple results
func swap(x, y string) (string, string) {
  return  y, x
}

// named return value
func split(sum int) (x, y int){
  x = sum * 4 / 9
  y = sum - x
  return
}
