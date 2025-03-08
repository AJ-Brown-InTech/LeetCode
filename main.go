package main

import (
    "log"
  "LeetCode/twosum"
)

func main() {
    if test, _ := twosum.Setup(); len(test) > 0 {
        log.Println(test)
        return
    }
}    


