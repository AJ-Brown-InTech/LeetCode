package main

import (
    "log"
    "github.com/AJ-Brown-InTech/LeetCode/twosum" 
)

func main() {
    if err := twosum.Setup(); len(err) > 0 {
        log.Println(err)
        return
    }
}    


