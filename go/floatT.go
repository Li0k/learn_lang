package main

import (
  "fmt"
)

func main()  {
  // var x float64 = 4.0
  var result = 8.001
  fmt.Printf("%T(%v)\n",result,result)
}

func fun(x float64) float64 {
  y := 3.0 * x - 4.0
  return y
}
