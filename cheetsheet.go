package main

import (
  "fmt"
  "math"
  "math/cmplx"
)

func main() {
  fmt.Println("Now you have %g problems.\n", math.Sqrt(7))
  fmt.Println(add(42, 13))
  a, b := swap("hello", "world")
  fmt.Println(a, b)
  fmt.Println(split(17))
  variables()
  basic_types()
  zero_values()
  type_conversions()
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

// Variables
func variables() {
  // initialize
  var i, j int = 1, 2

  // short variable declarations
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i, j, k, c, python, java)
}

// Basic Types
func basic_types(){
  var (
    ToBe bool = false
    MaxInt uint64 = 1<<64 - 1
    z complex128 = cmplx.Sqrt(-5 + 12i)
  )
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Zero values
func zero_values(){
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// Type conversions
func type_conversions(){
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)
}
