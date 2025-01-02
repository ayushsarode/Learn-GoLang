package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  msgChan := make(chan string)

  go func() {
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    msgChan <- "hello" // Write data to the channel
    msgChan <- "world" // Write data to the channel
  }()

  msg1 := <- msgChan
  msg2 := <- msgChan

  fmt.Println(msg1, msg2)
}
