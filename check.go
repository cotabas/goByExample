package main

import "fmt"

type counter struct {
  init, count int
}

func (c *counter) reset() int {
  c.count = c.init
  return c.init
}

func (c *counter) increment() int {
  c.count++
  return c.count
}

func (c *counter) decrement() int {
  c.count--
  return c.count
}

func main() {

  tester := counter{init: 5, count: 5}
 
  fmt.Println("start: ", tester)

  fmt.Println("++", tester.increment())
  fmt.Println("++", tester.increment())
  fmt.Println("++", tester.increment())
  fmt.Println("++", tester.increment())
  fmt.Println("++", tester.increment())
  fmt.Println("--", tester.decrement())
  fmt.Println("reset", tester.reset())
  fmt.Println("--", tester.decrement())
  fmt.Println("--", tester.decrement())
  fmt.Println("--", tester.decrement())

}
