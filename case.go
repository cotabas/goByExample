package main

import (
  "fmt"
  "time"
)

func main() {

  i := 1
  word := ""
  t := time.Now()

  switch i {
  case 1:
    word = "one"
  case 2:
    word = "two"
    case 3: 
    word = "three"
  }

  fmt.Println("write", i, "as", word)

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  default:
    fmt.Println("regular day..")
  }

  switch {
  case t.Hour() < 12:
    fmt.Println("pre-noon")
  default:
    fmt.Println("post-noon")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("boolio")
    case int:
      fmt.Println("inty")
    default:
      fmt.Printf("What's a %T?\n", t)
    }
  }

  whatAmI(true)
  whatAmI(3)
  whatAmI("string")
  whatAmI(3.14)
}


