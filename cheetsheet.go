package main

import (
  "fmt"
  "math"
  "math/cmplx"
  "runtime"
  "time"
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
  type_inference()
  constants()
  numeric_constants()
  for_statement()
  for_while()
  fmt.Println(sqrt(2), sqrt(-4))
  fmt.Println(pow(3,2,10), pow(3,3,20))
  fmt.Println(pow2(3,2,10), pow2(3,3,20))
  switch_statement()
  switch_eval()
  defer_statement()
  stack_defers()
  pointers()
  struct_statement()
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

// Type inference
func type_inference(){
  v := 42 // -> int
  fmt.Printf("v is of type %T\n", v)
}

// Constants
func constants(){
  // you can't use ":=".
  // only character, string, boolean, numeric
  const Pi = 3.14
  const World = "世界"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}

// Numeric Constants
func numeric_constants(){
  const(
    Big = 1 << 100
    Small = Big >> 99
  )

  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

// for
func for_statement(){
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}

// for is Go's while
func for_while(){
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}

// forever
// don't call this method
func forever(){
  for {
  }
}

// if
func sqrt(x float64) string{
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

// if with a short statement
func pow(x, n, lim float64) float64 {
  if v:= math.Pow(x, n); v < lim {
    return v
  }
  return lim
}

// if and else
func pow2(x,n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't v here, though
  return lim
}

// switch
func switch_statement() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd,
    // plan9 windows...
    fmt.Printf("%s.\n", os)
  }
}

// switch evaluation order
func switch_eval(){
  fmt.Println("When's Saturday?")
  today := time.Now().Weekday()
  switch time.Saturday {
  case today + 0:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 2:
    fmt.Println("In tow days.")
  default:
    fmt.Println("Too far away.")
  }
}

// defer
func defer_statement(){
  defer fmt.Println("world")
  fmt.Println("hello")
}

// stacking defers
func stack_defers() {
  fmt.Println("counting")

  for i:=0; i < 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("done")
}

// Pointers
func pointers() {
  i,j := 42, 2701

  p := &i         // point to i
  fmt.Println(*p) // read i through the pointer
  *p = 21         // set i through the pointer
  fmt.Println(i)  // see the new value of i

  p = &j          // point to j
  *p = *p / 37    // divide j through the pointer
  fmt.Println(j)  // see the new value of j
}

// structs
type Vertex struct {
  X int
  Y int
}

func struct_statement(){
  fmt.Println(Vertex{1,2})
}
