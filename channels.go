package main

import "fmt"

func main() {
  
  messages := make(chan string)

  go func() { messages <- "pong" }()

  msg := <-messages

  fmt.Println(msg)
}
