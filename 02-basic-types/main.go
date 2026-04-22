package main

import (
	"fmt"
	"goland/02-basic-types/logger"
	"os"
)

func main() {
  var sum float64
  var n int

  for {
  	var val float64

  	_, err := logger.LoggingError(&val)

    if err != nil {
    	break
    }

    sum += val
    n++
  }

  if n == 0 {
  	fmt.Fprint(os.Stderr, "no values")
  	os.Exit(1)
  }

  fmt.Println("The average is", sum/float64(n))
}
