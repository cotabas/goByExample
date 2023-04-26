package main

import "fmt"

func main() {

  m := make(map[int]string)

  m[1] = "this makes"
  m[2] = "more sense"

  fmt.Println("map:", m)


  v1 := m[1]
  fmt.Println("v1:", v1)

  v3 := m[3]
  fmt.Println("v3:", v3)

  fmt.Println("len:", len(m))


  delete(m, 2)
  fmt.Println("map:", m)

  _, prs := m[2]
  fmt.Println("prs:", prs)

  n := map[string]int{"foot": 1, "barf": 2}
  fmt.Println("n map:", n)
}

