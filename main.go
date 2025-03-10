package main

import (
    "log"
  "LeetCode/problem1"
)

func main() {
    if test, _ := problem1.Setup(); len(test) > 0 {
        log.Println(test)
        return
    }
}    


